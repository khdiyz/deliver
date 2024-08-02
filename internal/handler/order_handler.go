package handler

import (
	"deliver/internal/handler/response"
	"deliver/internal/models"
	"deliver/pkg/helper"
	"deliver/pkg/validator"
	"errors"

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
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/orders/{id}/receive-courier [post]
// @Security ApiKeyAuth
func (h *Handler) receiveOrderCourier(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	orderId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err = h.services.Order.ReceiveOrderCourier(orderId, userId); err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, gin.H{
		"status": "order received by courier",
	}, nil)
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

// Get List Order History
// @Description Get List Order History
// @Summary Get List Order History Customer
// @Tags Order
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/orders/history [get]
// @Security ApiKeyAuth
func (h *Handler) getOrderHistory(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	filters := make(map[string]interface{})
	filters["user-id"] = userId

	orders, err := h.services.Order.GetList(&pagination, filters)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, orders, &pagination)
}

// Finish Order
// @Description Finish Order
// @Summary Finish Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int64 true "Order Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/orders/{id}/finish-courier [post]
// @Security ApiKeyAuth
func (h *Handler) finishOrderCourier(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	orderId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err = h.services.Order.FinishOrderCourier(orderId, userId); err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, gin.H{
		"status": "order delivered by courier",
	}, nil)
}

// Payment Collect Order
// @Description Payment Collect Order
// @Summary Payment Collect Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int64 true "Order Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/orders/{id}/payment-collect [post]
// @Security ApiKeyAuth
func (h *Handler) paymentCollectOrderCourier(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	orderId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err = h.services.Order.PaymentCollectedOrderCourier(orderId, userId); err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, gin.H{
		"status": "payment collect by courier",
	}, nil)
}
