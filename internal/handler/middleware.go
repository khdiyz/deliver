package handler

import (
	"deliver/internal/constants"
	"deliver/internal/handler/response"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	UserCtx             = "user_id"
	RoleCtx             = "role_name"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(AuthorizationHeader)
	if header == "" {
		response.AbortResponse(c, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		response.AbortResponse(c, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		response.AbortResponse(c, "token is empty")
		return
	}

	claims, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		response.AbortResponse(c, err.Error())
		return
	}

	if claims.Type != "access" {
		response.AbortResponse(c, "invalid token type")
		return
	}

	c.Set(UserCtx, claims.UserId)
	c.Set(RoleCtx, claims.RoleName)
	c.Next()
}

func isAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(RoleCtx)
		if !exists || role != constants.RoleAdmin {
			response.AbortResponse(c, "permission denied")
			return
		}

		c.Next()
	}
}

func isCourierMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(RoleCtx)
		if !exists || role != constants.RoleCourier {
			response.AbortResponse(c, "permission denied")
			return
		}

		c.Next()
	}
}

func isCustomerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(RoleCtx)
		if !exists || role != constants.RoleCustomer {
			response.AbortResponse(c, "permission denied")
			return
		}

		c.Next()
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,Access-Control-Request-Method, Access-Control-Request-Headers")
		ctx.Header("Access-Control-Max-Age", "3600")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH, HEAD")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
