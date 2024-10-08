package handler

import (
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/pkg/logger"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func listPagination(c *gin.Context) (pagination models.Pagination, err error) {
	page, err := getPageQuery(c)
	if err != nil {
		logger.GetLogger().Error(err)
		return pagination, err
	}
	pageSize, err := getPageSizeQuery(c)
	if err != nil {
		logger.GetLogger().Error(err)
		return pagination, err
	}
	offset, limit := calculatePagination(page, pageSize)
	pagination.Limit = limit
	pagination.Offset = offset
	pagination.Page = page
	pagination.PageSize = pageSize
	return pagination, nil
}

func getPageQuery(c *gin.Context) (offset int64, err error) {
	offsetStr := c.DefaultQuery("page", constants.DefaultPage)
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error while parsing query: %v", err.Error())
	}
	if offset < 0 {
		return 0, fmt.Errorf("page should be unsigned")
	}
	return offset, nil
}

func getPageSizeQuery(c *gin.Context) (limit int64, err error) {
	limitStr := c.DefaultQuery("pageSize", constants.DefaultPageSize)
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error while parsing query: %v", err.Error())
	}
	if limit < 0 {
		return 0, fmt.Errorf("pageSize should be unsigned")
	}
	return limit, nil
}

func calculatePagination(page, pageSize int64) (offset, limit int64) {
	if page < 0 {
		page = 1
	}
	offset = (page - 1) * pageSize
	limit = pageSize
	return offset, limit
}

func getNullInt64Param(c *gin.Context, paramName string) (int64, error) {
	paramData := c.Param(paramName)

	if paramData != "" {
		paramValue, err := strconv.ParseInt(paramData, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid param: %s", paramData)
		}

		return paramValue, nil
	}

	return 0, errors.New("param required")
}

func getNullStringQuery(c *gin.Context, queryName string) string {
	queryData := c.Query(queryName)
	queryData = strings.Trim(queryData, " ")
	return queryData
}

func getUserId(ctx *gin.Context) (int64, error) {
	id, ok := ctx.Get("user_id")
	if !ok {
		return 0, constants.ErrInvalidUserId
	}
	userId, ok := id.(int64)
	if !ok {
		return 0, constants.ErrInvalidUserId
	}

	return userId, nil
}
