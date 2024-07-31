package handler

import (
	"deliver/internal/handler/response"
	"deliver/internal/models"
	"deliver/pkg/validator"

	"github.com/gin-gonic/gin"
)

const (
	idQuery = "id"
)

// Create Category
// @Description Create Category
// @Summary Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param create body models.CategoryCreateRequest true "Create Category"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/categories [post]
// @Security ApiKeyAuth
func (h *Handler) createCategory(c *gin.Context) {
	var (
		err   error
		input models.CategoryCreateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	id, err := h.services.Category.Create(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.Created, id, nil)
}

// Get List Category
// @Description Get List Category
// @Summary Get List Category
// @Tags Category
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/categories [get]
// @Security ApiKeyAuth
func (h *Handler) getListCategory(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	categories, err := h.services.Category.GetList(&pagination)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, categories, &pagination)
}

// Get Category By Id
// @Description Get Category By Id
// @Summary Get Category By Id
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int64 true "Category Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/categories/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getCategoryById(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	category, err := h.services.Category.GetById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, category, nil)
}

// Update Category
// @Description Update Category
// @Summary Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int64 true "Category Id"
// @Param update body models.CategoryUpdateRequest true "Update Category"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/categories/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateCategory(c *gin.Context) {
	var input models.CategoryUpdateRequest

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

	err = h.services.Category.Update(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}

// Delete Category
// @Description Delete Category
// @Summary Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int64 true "Category Id"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/categories/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteCategory(c *gin.Context) {
	id, err := getNullInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	err = h.services.Category.DeleteById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}
