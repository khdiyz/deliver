package handler

import (
	"deliver/internal/handler/response"
	"deliver/internal/models"
	"deliver/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create Attribute
// @Description Create Attribute
// @Summary Create Attribute
// @Tags Attribute
// @Accept json
// @Produce json
// @Param create body models.AttributeCreateRequest true "Create Attribute"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/attributes [post]
// @Security ApiKeyAuth
func (h *Handler) createAttribute(c *gin.Context) {
	var (
		err   error
		input models.AttributeCreateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	id, err := h.services.Attribute.Create(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.Created, id, nil)
}

// Get List Attribute
// @Description Get List Attribute
// @Summary Get List Attribute
// @Tags Attribute
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/attributes [get]
// @Security ApiKeyAuth
func (h *Handler) getListAttribute(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	attributes, err := h.services.Attribute.GetList(&pagination)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, attributes, &pagination)
}

// Get Attribute By Id
// @Description Get Attribute By Id
// @Summary Get Attribute By Id
// @Tags Attribute
// @Accept json
// @Produce json
// @Param id path int64 true "Attribute Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/attributes/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getAttributeById(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	attribute, err := h.services.Attribute.GetById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, attribute, nil)
}

// Update Attribute
// @Description Update Attribute
// @Summary Update Attribute
// @Tags Attribute
// @Accept json
// @Produce json
// @Param id path int64 true "Attribute Id"
// @Param update body models.AttributeUpdateRequest true "Update Attribute"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/attributes/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateAttribute(c *gin.Context) {
	var input models.AttributeUpdateRequest

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

	err = h.services.Attribute.Update(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}

// Delete Attribute
// @Description Delete Attribute
// @Summary Delete Attribute
// @Tags Attribute
// @Accept json
// @Produce json
// @Param id path int64 true "Attribute Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/attributes/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteAttribute(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	err = h.services.Attribute.DeleteById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}
