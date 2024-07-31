package service

import (
	"deliver/internal/repository"
	"deliver/internal/storage"
	"deliver/models"
	"deliver/pkg/logger"
	"io"
	"time"
)

type Service struct {
	Authorization
	User
	Minio
	Role
}

func NewService(repo repository.Repository, storage storage.Storage, log logger.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(repo, log),
		User:          NewUserService(repo, log),
		Minio:         NewMinioService(storage, log),
		Role:          NewRoleService(repo, log),
	}
}

type Authorization interface {
	CreateToken(user models.User, tokenType string, expiresAt time.Time) (*models.Token, error)
	GenerateTokens(user models.User) (*models.Token, *models.Token, error)
	ParseToken(token string) (*jwtCustomClaim, error)
	Login(input models.LoginRequest) (*models.Token, *models.Token, error)
}

type User interface {
	Create(input models.UserCreateRequest) (int64, error)
	GetByEmail(email string) (models.User, error)
	GetById(id int64) (models.User, error)
}

type Minio interface {
	UploadImage(image io.Reader, imageSize int64, contextType string) (storage.File, error)
	UploadDoc(doc io.Reader, docSize int64, contextType string) (storage.File, error)
	UploadExcel(doc io.Reader, docSize int64, contextType string) (storage.File, error)
}

type Role interface {
	GetList(pagination *models.Pagination) ([]models.Role, error)
}
