package handler

import (
	"deliver/internal/handler/response"

	"github.com/gin-gonic/gin"
)

// GetListRole
// @Description Get List Role
// @Summary Get List Role
// @Tags Role
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/roles [get]
// @Security ApiKeyAuth
func (h *Handler) getListRole(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	roles, err := h.services.Role.GetList(&pagination)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, roles, &pagination)
}
