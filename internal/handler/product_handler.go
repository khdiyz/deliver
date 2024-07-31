package handler

import (
	"deliver/internal/handler/response"
	"deliver/models"
	"deliver/pkg/validator"

	"github.com/gin-gonic/gin"
)

const (
	attributeIdParam = "attribute-id"
)

// Create Product
// @Description Create Product
// @Summary Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param create body models.ProductCreateRequest true "Create Product"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/products [post]
// @Security ApiKeyAuth
func (h *Handler) createProduct(c *gin.Context) {
	var (
		err   error
		input models.ProductCreateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	id, err := h.services.Product.Create(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.Created, id, nil)
}

// Get List Product
// @Description Get List Product
// @Summary Get List Product
// @Tags Product
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/products [get]
// @Security ApiKeyAuth
func (h *Handler) getListProduct(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	products, err := h.services.Product.GetList(&pagination)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, products, &pagination)
}

// Get Product By Id
// @Description Get Product By Id
// @Summary Get Product By Id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int64 true "Product Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/products/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getProductById(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	product, err := h.services.Product.GetById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, product, nil)
}

// Update Product
// @Description Update Product
// @Summary Update Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int64 true "Product Id"
// @Param update body models.ProductUpdateRequest true "Update Product"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/products/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateProduct(c *gin.Context) {
	var input models.ProductUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}
	input.Id = id

	err = h.services.Product.Update(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}

// Delete Product
// @Description Delete Product
// @Summary Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int64 true "Product Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/products/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteProduct(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	err = h.services.Product.DeleteById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}

// Add Attribute to Product
// @Description Add Attribute to Product
// @Summary Add Attribute to Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int64 true "Product Id"
// @Param attribute-id path int64 true "Attribute Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/products/{id}/add/{attribute-id} [post]
// @Security ApiKeyAuth
func (h *Handler) addAttributeToProduct(c *gin.Context) {
	productId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	attributeId, err := getNullInt64Param(c, attributeIdParam)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	err = h.services.Product.AddAttributeToProduct(productId, attributeId)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}

// Remove Attribute from Product
// @Description Remove Attribute from Product
// @Summary Remove Attribute from Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int64 true "Product Id"
// @Param attribute-id path int64 true "Attribute Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/products/{id}/remove/{attribute-id} [delete]
// @Security ApiKeyAuth
func (h *Handler) removeAttributeFromProduct(c *gin.Context) {
	productId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	attributeId, err := getNullInt64Param(c, attributeIdParam)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	err = h.services.Product.RemoveAttributeFromProduct(productId, attributeId)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}
