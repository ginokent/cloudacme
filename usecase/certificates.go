package usecase

import (
	"context"
	"crypto"
	"crypto/tls"
	"crypto/x509"

	"github.com/newtstat/cloudacme/contexts"
	"github.com/newtstat/cloudacme/repository"
	"github.com/newtstat/cloudacme/trace"
	"github.com/newtstat/nits.go"
	"golang.org/x/xerrors"
)

type CertificatesUseCase interface {
	IssueCertificate(ctx context.Context, privateKey crypto.PrivateKey, privateKeyRenewed bool, privateKeyVaultResource string, certificateVaultResource string, thresholdOfDaysToExpire int64, domains []string) (privateKeyVaultVersionResource, certificateVaultVersionResource string, err error)
}

var _ CertificatesUseCase = (*certificatesUseCase)(nil)

type certificatesUseCase struct {
	vaultRepo       repository.VaultRepository
	letsencryptRepo repository.LetsEncryptRepository
}

// nolint: ireturn
func NewCertificatesUseCase(certificatesRepo repository.VaultRepository, letsencryptRepo repository.LetsEncryptRepository) CertificatesUseCase {
	return CertificatesUseCase(&certificatesUseCase{
		vaultRepo:       certificatesRepo,
		letsencryptRepo: letsencryptRepo,
	})
}

func (uc *certificatesUseCase) IssueCertificate(
	ctx context.Context,
	privateKey crypto.PrivateKey,
	privateKeyRenewed bool,
	privateKeyVaultResource string,
	certificateVaultResource string,
	thresholdOfDaysToExpire int64,
	domains []string,
) (
	privateKeyVaultVersionResource,
	certificateVaultVersionResource string,
	err error,
) {
	return uc.issueCertificate(
		ctx,
		privateKey,
		privateKeyRenewed,
		privateKeyVaultResource,
		certificateVaultResource,
		thresholdOfDaysToExpire,
		domains,
		nits.X509.CheckCertificatePEM,
		nits.X509.ParsePKCSXPrivateKeyPEM,
		tls.X509KeyPair,
	)
}

// nolint: cyclop, funlen
func (uc *certificatesUseCase) issueCertificate(
	ctx context.Context,
	privateKey crypto.PrivateKey,
	privateKeyRenewed bool,
	privateKeyVaultResource string,
	certificateVaultResource string,
	thresholdOfDaysToExpire int64,
	domains []string,
	checkCertificatePEMFunc func(pemData []byte) (notyet bool, daysToStart int64, expired bool, daysToExpire int64, err error),
	parsePKCSXPrivateKeyPEMFunc func(pemData []byte) (crypto.PrivateKey, error),
	tls_X509KeyPair func(certPEMBlock []byte, keyPEMBlock []byte) (tls.Certificate, error), // nolint: revive
) (
	privateKeyVaultVersionResource,
	certificateVaultVersionResource string,
	err error,
) {
	ctx, span := trace.Start(ctx, "(*usecase.certificatesUseCase).Issue")
	defer span.End()

	l := contexts.GetLogger(ctx)

	privateKeyExists, privateKeyVaultVersionResource, privateKeyErr := uc.vaultRepo.GetVaultVersionIfExists(ctx, privateKeyVaultResource+"/versions/latest")
	certificateExists, certificateVaultVersionResource, certificatePEM, certificateErr := uc.vaultRepo.GetVaultVersionDataIfExists(ctx, certificateVaultResource+"/versions/latest")
	if privateKeyErr != nil || certificateErr != nil {
		return "", "", xerrors.Errorf("(*usecase.certificatesUseCase).vaultRepo.GetVaultVersionDataIfExists: privateKeyErr=%v, certificateErr=%w", privateKeyErr, certificateErr)
	}

	var keyPairIsBroken bool
	// nolint: godox
	// TODO
	privateKeyPEM, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		l.E().Error(xerrors.Errorf("🚨 private key is broken. tls.X509KeyPair: %w", err))
		keyPairIsBroken = true
	} else {
		if _, err := tls_X509KeyPair(certificatePEM, privateKeyPEM); err != nil {
			l.E().Error(xerrors.Errorf("🚨 certificate key pair is broken. tls.X509KeyPair: %w", err))
			keyPairIsBroken = true
		}
	}

	if !privateKeyRenewed && // NOTE: If renewPrivateKey, skip checking certificate and force to renew certificate
		privateKeyExists && certificateExists &&
		!keyPairIsBroken {
		l.Info("🔬 checking certificate...")

		var notyet bool
		var expired bool
		var daysToExpire int64
		if err := trace.StartFunc(ctx, "nits.X509.CheckCertificatePEM")(func(child context.Context) (err error) {
			notyet, _, expired, daysToExpire, err = checkCertificatePEMFunc(certificatePEM)
			return
		}); err != nil {
			l.E().Error(xerrors.Errorf("🚨 certificate (%s) is broken. nits.X509.CheckCertificatePEM: %w", certificateVaultVersionResource, err))
		} else if !notyet && !expired && daysToExpire > thresholdOfDaysToExpire {
			l.F().Infof("✅ there is still time (%d days) for current certificate to expire. It will not be renewed", daysToExpire)
			return privateKeyVaultVersionResource, certificateVaultVersionResource, nil // early return
		}

		l.F().Infof("❗️ current certificate has expired or is due to expire in less than %d days. Renew the certificate", thresholdOfDaysToExpire)
	}

	l.Info("🪪 generate certificate...")

	if err := uc.vaultRepo.CreateVaultIfNotExists(ctx, certificateVaultResource); err != nil {
		return "", "", xerrors.Errorf("(*usecase.certificatesUseCase).vaultRepo.CreateVaultIfNotExists: %w", err)
	}

	privateKeyPEM, certificatePEM, _, _, err = uc.letsencryptRepo.IssueCertificate(ctx, privateKey, domains)
	if err != nil {
		return "", "", xerrors.Errorf("(*usecase.certificatesUseCase).letsencryptRepo.IssueCertificate: %w", err)
	}

	l.Info("🪪 generated certificate")

	if privateKeyRenewed {
		privateKeyVaultVersionResource, err = uc.vaultRepo.AddVaultVersion(ctx, privateKeyVaultResource, privateKeyPEM)
		if err != nil {
			return "", "", xerrors.Errorf("(*usecase.certificatesUseCase).vaultRepo.AddVaultVersion: %w", err)
		}
	}

	certificateVaultVersionResource, err = uc.vaultRepo.AddVaultVersion(ctx, certificateVaultResource, certificatePEM)
	if err != nil {
		return "", "", xerrors.Errorf("(*usecase.certificatesUseCase).vaultRepo.AddVaultVersion: %w", err)
	}

	l.Info("✅ save certificate")

	return privateKeyVaultVersionResource, certificateVaultVersionResource, nil
}
