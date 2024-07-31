package handler

import (
	"deliver/internal/handler/response"
	"deliver/models"
	"deliver/pkg/validator"
	"errors"

	"github.com/gin-gonic/gin"
)

// Login
// @Description Login User
// @Summary Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input models.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	accessToken, refreshToken, err := h.services.Authorization.Login(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, gin.H{
		"accessToken":  accessToken.Token,
		"refreshToken": refreshToken.Token,
	}, nil)
}

// Refresh token
// @Description Refresh Token
// @Summary Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param token body models.RefreshRequest true "Refresh Token"
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 404 {object} models.BaseResponse
// @Failure 500 {object} models.BaseResponse
// @Router /api/v1/auth/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	var (
		err   error
		input models.RefreshRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err = validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	claims, err := h.services.Authorization.ParseToken(input.Token)
	if err != nil {
		response.ErrorResponse(c, response.Unauthorized, err)
		return
	}

	if claims.Type != "refresh" {
		response.ErrorResponse(c, response.BadRequest, errors.New("token type must be refresh"))
		return
	}

	user, err := h.services.User.GetById(claims.UserId)
	if err != nil {
		response.FromError(c, err)
		return
	}

	accessToken, refreshToken, err := h.services.Authorization.GenerateTokens(user)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, gin.H{
		"accessToken":  accessToken.Token,
		"refreshToken": refreshToken.Token,
	}, nil)
}
