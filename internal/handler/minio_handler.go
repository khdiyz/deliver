package handler

import (
	"deliver/internal/handler/response"
	"deliver/pkg/logger"
	"errors"
	"io"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

const (
	contentType     = "Content-Type"
	filePath        = "file-path"
	jpgContentType  = "image/jpg"
	pngContentType  = "image/png"
	jpegContentType = "image/jpeg"
	webpContentType = "image/webp"
	xlsxContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	xlsContentType  = "application/vnd.ms-excel"
	docContentType  = "application/msword"
	pdfContentType  = "application/pdf"
	docxContentType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
)

// UploadImage
// @Description Upload Image
// @Tags Minio
// @Accept       json
// @Produce application/octet-stream
// @Produce image/png
// @Produce image/jpeg
// @Produce image/jpg
// @Param file formData file true "file"
// @Accept multipart/form-data
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/minio/upload-image [post]
// @Security ApiKeyAuth
func (h *Handler) uploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	imageContentType := file.Header[contentType][0]
	if imageContentType != jpegContentType && imageContentType != jpgContentType && imageContentType != pngContentType && imageContentType != webpContentType {
		response.ErrorResponse(c, response.BadRequest, errors.New("invalid file format"))
		return
	}

	var fileIO io.Reader
	fileMultipart, err := file.Open()
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	fileIO = fileMultipart
	imageFile, err := h.services.Minio.UploadImage(fileIO, file.Size, imageContentType)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	defer func(fileMultipart multipart.File) {
		err := fileMultipart.Close()
		if err != nil {
			logger.GetLogger().Error(err)
		}
	}(fileMultipart)

	response.SuccessResponse(c, response.OK, imageFile, nil)
}
