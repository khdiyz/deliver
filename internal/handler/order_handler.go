package handler

import (
	"deliver/internal/handler/response"
	"deliver/internal/models"
	"deliver/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create Order
// @Description Create Order
// @Summary Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param create body models.OrderCreateRequest true "Create Order"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/orders [post]
// @Security ApiKeyAuth
func (h *Handler) createOrder(c *gin.Context) {
	var (
		err   error
		input models.OrderCreateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	input.RecieverId, err = getUserId(c)
	if err != nil {
		response.ErrorResponse(c, response.NotFound, err)
		return
	}

	orderId, err := h.services.Order.CreateOrder(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, orderId, nil)
}
