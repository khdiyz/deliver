package handler

import (
	"deliver/internal/handler/response"
	"deliver/internal/models"
	"deliver/pkg/helper"
	"deliver/pkg/validator"
	"errors"
	"net/http"

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

// Recieve Order
// @Description Recieve Order
// @Summary Recieve Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int64 true "Order Id"
// @Param recieve body models.OrderCourierRequest true "Recieve Order"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/orders/{id}/receive-courier [post]
// @Security ApiKeyAuth
func (h *Handler) receiveOrderCourier(c *gin.Context) {
	var input models.OrderCourierRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	orderId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	err = h.services.Order.ReceiveOrderCourier(orderId, input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "order received by courier"})
}

// Get Order By Id
// @Description Get Order By Id
// @Summary Get Order By Id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int64 true "Order Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/orders/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getOrderById(c *gin.Context) {
	orderId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	order, err := h.services.Order.GetById(orderId)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, order, nil)
}

// Get List Order
// @Description Get List Order
// @Summary Get List Order
// @Tags Order
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Param status query string false "status" Enums(picked_up, on_delivery, delivered, payment_collected)
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/orders [get]
// @Security ApiKeyAuth
func (h *Handler) getListOrder(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	filters := make(map[string]interface{})
	status := getNullStringQuery(c, "status")
	if status != "" {
		if !helper.IsValidOrderStatus(status) {
			response.ErrorResponse(c, response.BadRequest, errors.New("invalid status name"))
			return
		}
		filters["status"] = status
	}

	orders, err := h.services.Order.GetList(&pagination, filters)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, orders, &pagination)
}
