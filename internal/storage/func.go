package storage

import (
	"context"
	"deliver/config"
	"io"
	"strings"

	"github.com/minio/minio-go/v7"
)

var (
	docContentType  = "msword"
	docxContentType = "vnd.openxmlformats-officedocument.wordprocessingml.document"

	xlsxContentType = "vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	xlsContentType  = "vnd.ms-excel"
)

// getFileExtension extracts the file extension from the content type.
func getFileExtension(contentType string) string {
	parts := strings.Split(contentType, "/")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

// getFileExtensionForDoc determines the file extension for the document based on the content type.
func getFileExtensionForDoc(contentType string) string {
	switch {
	case strings.Contains(contentType, docContentType):
		return "doc"
	case strings.Contains(contentType, docxContentType):
		return "docx"
	default:
		return "docx" // Default to docx if content type is not recognized
	}
}

// getFileExtensionForExcel determines the file extension for the Excel file based on the content type.
func getFileExtensionForExcel(contentType string) string {
	switch {
	case strings.Contains(contentType, xlsContentType):
		return "xls"
	case strings.Contains(contentType, xlsxContentType):
		return "xlsx"
	default:
		return "xlsx" // Default to xlsx if content type is not recognized
	}
}

// uploadToMinio uploads the file to MinIO.
func (um *UploadMinio) uploadToMinio(fileName string, file io.Reader, fileSize int64, contentType string) error {
	_, err := um.minio.PutObject(
		context.Background(),
		um.cfg.MinioBucketName,
		fileName,
		file,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	return err
}

func GenerateLink(cfg config.Config, fileName string) string {
	if fileName == "" {
		return ""
	}

	return cfg.MinioEndpoint + "/" + cfg.MinioBucketName + "/" + fileName
}
