package setup

import (
	"deliver/config"
	"deliver/pkg/logger"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func SetupMinioConnection(cfg *config.Config, log *logger.Logger) (*minio.Client, error) {
	// Initialize minio client object.
	minioClient, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyId, cfg.MinioSecretKey, ""),
		Secure: cfg.MinioUseSSL,
	})
	if err != nil {
		log.Fatal(err)
	}

	return minioClient, err
}
