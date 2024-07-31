package storage

import (
	"deliver/config"
	"deliver/pkg/logger"
	"io"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type File struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type UploadMinio struct {
	minio  *minio.Client
	cfg    *config.Config
	logger *logger.Logger
}

func NewUploadMinio(minio *minio.Client, cfg *config.Config, logger *logger.Logger) *UploadMinio {
	return &UploadMinio{
		minio:  minio,
		cfg:    cfg,
		logger: logger,
	}
}

// UploadImage uploads an image file to MinIO and returns the generated file name or an error.
func (um *UploadMinio) UploadImage(image io.Reader, imageSize int64, contentType string) (file File, err error) {
	fileName := uuid.NewString()

	// Extract the file extension from the content type
	fileExtension := getFileExtension(contentType)

	// Construct the image name
	file.Name = fileName + "." + fileExtension

	// Upload the file to MinIO
	if err := um.uploadToMinio(file.Name, image, imageSize, contentType); err != nil {
		um.logger.Error("Internal Server Error: ", err)
		return file, err
	}

	// Generate the link for the uploaded image
	file.Link = GenerateLink(um.cfg, file.Name)

	return file, nil
}

func (um *UploadMinio) UploadDoc(doc io.Reader, docSize int64, contentType string) (file File, err error) {
	// Generate a new UUID for the file name
	fileName := uuid.NewString()

	// Determine the file extension based on the content type
	fileExtension := getFileExtensionForDoc(contentType)

	// Construct the document file name
	file.Name = fileName + "." + fileExtension

	// Upload the document file using the uploadToMinio method
	if err := um.uploadToMinio(file.Name, doc, docSize, contentType); err != nil {
		um.logger.Error("Internal Server Error: ", err)
		return file, err
	}

	// Generate the link for the uploaded document
	file.Link = GenerateLink(um.cfg, file.Name)

	return file, nil
}

func (um *UploadMinio) UploadExcel(excel io.Reader, excelSize int64, contentType string) (file File, err error) {
	// Generate a new UUID for the file name
	fileName := uuid.NewString()

	// Determine the file extension based on the content type
	fileExtension := getFileExtensionForExcel(contentType)

	// Construct the Excel file name
	file.Name = fileName + "." + fileExtension

	// Upload the Excel file using the uploadToMinio method
	if err := um.uploadToMinio(file.Name, excel, excelSize, contentType); err != nil {
		um.logger.Error("Internal Server Error: ", err)
		return file, err
	}

	// Generate the link for the uploaded Excel file
	file.Link = GenerateLink(um.cfg, file.Name)

	return file, nil
}
