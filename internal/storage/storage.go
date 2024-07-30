package storage

import (
	"deliver/config"
	"deliver/pkg/logger"
	"io"

	"github.com/minio/minio-go/v7"
)

type Storage struct {
	UploadStorage
}

func NewStorage(minio *minio.Client, cfg *config.Config, log *logger.Logger) *Storage {
	return &Storage{
		UploadStorage: NewUploadMinio(minio, cfg, log),
	}
}

type UploadStorage interface {
	UploadImage(image io.Reader, imageSize int64, contentType string) (file File, err error)
	UploadDoc(doc io.Reader, docSize int64, contentType string) (file File, err error)
	UploadExcel(excel io.Reader, excelSize int64, contentType string) (file File, err error)
}
