package service

import (
	"deliver/internal/storage"
	"deliver/pkg/logger"
	"io"

	"google.golang.org/grpc/codes"
)

type MinioService struct {
	storage storage.Storage
	log     logger.Logger
}

func NewMinioService(storage storage.Storage, log logger.Logger) *MinioService {
	return &MinioService{storage: storage, log: log}
}

func (m *MinioService) UploadImage(image io.Reader, imageSize int64, contextType string) (storage.File, error) {
	file, err := m.storage.UploadStorage.UploadImage(image, imageSize, contextType)
	if err != nil {
		return storage.File{}, serviceError(err, codes.Internal)
	}

	return file, nil
}

func (m *MinioService) UploadDoc(doc io.Reader, docSize int64, contextType string) (storage.File, error) {
	file, err := m.storage.UploadStorage.UploadDoc(doc, docSize, contextType)
	if err != nil {
		return storage.File{}, serviceError(err, codes.Internal)
	}

	return file, nil
}

func (m *MinioService) UploadExcel(doc io.Reader, docSize int64, contextType string) (storage.File, error) {
	file, err := m.storage.UploadStorage.UploadExcel(doc, docSize, contextType)
	if err != nil {
		return storage.File{}, serviceError(err, codes.Internal)
	}

	return file, nil
}
