package service

import (
	"database/sql"
	"deliver/config"
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/internal/repository"
	"deliver/pkg/helper"
	"deliver/pkg/logger"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
)

type AuthService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewAuthService(repo repository.Repository, log logger.Logger) *AuthService {
	return &AuthService{
		repo: repo,
		log:  log,
	}
}

type jwtCustomClaim struct {
	jwt.StandardClaims
	UserId   int64  `json:"user_id"`
	RoleName string `json:"role"`
	Type     string `json:"type"`
}

func (s *AuthService) CreateToken(user models.User, tokenType string, expiresAt time.Time) (*models.Token, error) {
	claims := &jwtCustomClaim{
		UserId:   user.Id,
		RoleName: user.RoleName,
		Type:     tokenType,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return &models.Token{
		User:      user,
		Token:     token,
		Type:      tokenType,
		ExpiresAt: expiresAt,
	}, nil
}

func (s *AuthService) GenerateTokens(user models.User) (*models.Token, *models.Token, error) {
	accessExpiresAt := time.Now().Add(time.Duration(config.GetConfig().JWTAccessExpirationHours) * time.Hour)
	refreshExpiresAt := time.Now().Add(time.Duration(config.GetConfig().JWTRefreshExpirationDays) * time.Hour * 24)

	accessToken, err := s.CreateToken(user, constants.TokenTypeAccess, accessExpiresAt)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := s.CreateToken(user, constants.TokenTypeRefresh, refreshExpiresAt)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) ParseToken(token string) (*jwtCustomClaim, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*jwtCustomClaim)
	if !ok {
		return nil, errors.New("token claims are not of type *jwtCustomClaim")
	}

	return claims, nil
}

func (s *AuthService) Login(input models.LoginRequest) (*models.Token, *models.Token, error) {
	user, err := s.repo.User.GetByEmail(input.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, serviceError(errors.New("wrong username or password"), codes.Unauthenticated)
		}
		return nil, nil, serviceError(err, codes.Internal)
	}

	hashPassword, err := helper.GenerateHash(input.Password)
	if err != nil {
		return nil, nil, serviceError(err, codes.Internal)
	}

	if user.Password != hashPassword {
		return nil, nil, serviceError(errors.New("wrong username or password"), codes.Unauthenticated)
	}

	return s.GenerateTokens(user)
}

func (s *AuthService) SignUp(input models.SignUpRequest) (*models.Token, *models.Token, error) {
	_, err := s.repo.User.GetByEmail(input.Email)
	if err == nil {
		return nil, nil, serviceError(errors.New("user already exists with this email"), codes.InvalidArgument)
	} else if err != sql.ErrNoRows {
		return nil, nil, serviceError(err, codes.Internal)
	}

	input.Password, err = helper.GenerateHash(input.Password)
	if err != nil {
		return nil, nil, serviceError(err, codes.Internal)
	}

	if input.RoleName != constants.RoleCourier && input.RoleName != constants.RoleCustomer {
		return nil, nil, serviceError(errors.New("role name must be only: COURIER, CUSTOMER"), codes.InvalidArgument)
	}

	role, err := s.repo.Role.GetByName(input.RoleName)
	if err != nil {
		return nil, nil, serviceError(err, codes.Internal)
	}

	userId, err := s.repo.User.Create(models.UserCreateRequest{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
		RoleId:   role.Id,
	})
	if err != nil {
		return nil, nil, serviceError(err, codes.Internal)
	}

	return s.GenerateTokens(models.User{
		Id:       userId,
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
		RoleId:   role.Id,
		RoleName: role.Name,
	})
}
