package handler

import (
	"deliver/internal/handler/response"
	"deliver/models"
	"deliver/pkg/validator"

	"github.com/gin-gonic/gin"
)

const (
	optionIdParam    = "option-id"
	attributeIdQuery = "attributeId"
)

// Create Option
// @Description Create Option
// @Summary Create Option
// @Tags Option
// @Accept json
// @Produce json
// @Param id path int64 true "Attribute Id"
// @Param create body models.OptionCreateRequest true "Create Option"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/attributes/{id}/options [post]
// @Security ApiKeyAuth
func (h *Handler) createOption(c *gin.Context) {
	var (
		err   error
		input models.OptionCreateRequest
	)

	attributeId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}
	input.AttributeId = attributeId

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	id, err := h.services.Option.Create(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.Created, id, nil)
}

// Get List Option
// @Description Get List Option
// @Summary Get List Option
// @Tags Option
// @Accept json
// @Produce json
// @Param id path int64 true "Attribute Id"
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/attributes/{id}/options [get]
// @Security ApiKeyAuth
func (h *Handler) getListOption(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	filter := make(map[string]interface{})
	attributeId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}
	filter["attribute-id"] = attributeId

	options, err := h.services.Option.GetList(&pagination, filter)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, options, &pagination)
}

// Get Option By Id
// @Description Get Option By Id
// @Summary Get Option By Id
// @Tags Option
// @Accept json
// @Produce json
// @Param id path int64 true "Attribute Id"
// @Param option-id path int64 true "Option Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/attributes/{id}/options/{option-id} [get]
// @Security ApiKeyAuth
func (h *Handler) getOptionById(c *gin.Context) {
	attributeId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	optionId, err := getNullInt64Param(c, optionIdParam)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	option, err := h.services.Option.Get(attributeId, optionId)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, option, nil)
}

// Update Option
// @Description Update Option
// @Summary Update Option
// @Tags Option
// @Accept json
// @Produce json
// @Param id path int64 true "Attribute Id"
// @Param option-id path int64 true "Option Id"
// @Param update body models.OptionUpdateRequest true "Update Option"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/attributes/{id}/options/{option-id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateOption(c *gin.Context) {
	var input models.OptionUpdateRequest

	attributeId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	optionId, err := getNullInt64Param(c, optionIdParam)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}
	input.AttributeId = attributeId
	input.Id = optionId

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	err = h.services.Option.Update(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}

// Delete Option
// @Description Delete Option
// @Summary Delete Option
// @Tags Option
// @Accept json
// @Produce json
// @Param id path int64 true "Attribute Id"
// @Param option-id path int64 true "Option Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/attributes/{id}/options/{option-id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteOption(c *gin.Context) {
	attributeId, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	optionId, err := getNullInt64Param(c, optionIdParam)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	err = h.services.Option.Delete(attributeId, optionId)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}
