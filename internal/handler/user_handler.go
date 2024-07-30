package handler

import (
	"deliver/internal/handler/response"
	"deliver/models"
	"deliver/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create User
// @Description Create User
// @Summary Create User
// @Tags User
// @Accept json
// @Produce json
// @Param create body models.UserCreateRequest true "Create User"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/users [post]
// @Security ApiKeyAuth
func (h *Handler) createUser(c *gin.Context) {
	var (
		err   error
		input models.UserCreateRequest
	)
	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	id, err := h.services.User.Create(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.Created, id, nil)
}
