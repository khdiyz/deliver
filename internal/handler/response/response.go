package response

import (
	"deliver/internal/models"
	"errors"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SuccessResponse(c *gin.Context, status Status, data interface{}, pagination *models.Pagination) {
	c.JSON(status.Code, models.BaseResponse{
		Success:     true,
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
		Pagination:  pagination,
	})
}

func ErrorResponse(c *gin.Context, status Status, err error) {
	c.JSON(status.Code, models.BaseResponse{
		Success:      false,
		Status:       status.Status,
		Description:  status.Description,
		ErrorMessage: err.Error(),
	})
}

func AbortResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(Aborted.Code, models.BaseResponse{
		Success:      false,
		Status:       Aborted.Status,
		Description:  Aborted.Description,
		ErrorMessage: message,
	})
}

// CONVERT SERVICE ERROR TO HANDLER ERROR

func FromError(c *gin.Context, serviceError error) {
	st, _ := status.FromError(serviceError)
	err := st.Message()

	switch st.Code() {
	case codes.Internal:
		ErrorResponse(c, Internal, errors.New(err))
	case codes.NotFound:
		ErrorResponse(c, NotFound, errors.New(err))
	case codes.InvalidArgument:
		ErrorResponse(c, BadRequest, errors.New(err))
	case codes.Unavailable:
		ErrorResponse(c, Unavailable, errors.New(err))
	case codes.AlreadyExists:
		ErrorResponse(c, AlreadyExists, errors.New(err))
	case codes.Unauthenticated:
		ErrorResponse(c, Unauthorized, errors.New(err))
	}
}
