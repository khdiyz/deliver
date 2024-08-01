package service

import (
	"deliver/config"
	"deliver/internal/models"
	"deliver/internal/repository"
	"deliver/internal/storage"
	"deliver/pkg/logger"
	"io"
	"time"
)

type Service struct {
	Authorization
	User
	Minio
	Role
	Category
	Product
	Attribute
	Option
}

func NewService(repo repository.Repository, storage storage.Storage, log logger.Logger, cfg config.Config) *Service {
	return &Service{
		Authorization: NewAuthService(repo, log),
		User:          NewUserService(repo, log),
		Minio:         NewMinioService(storage, log),
		Role:          NewRoleService(repo, log),
		Category:      NewCategoryService(repo, log),
		Product:       NewProductService(repo, log, cfg),
		Attribute:     NewAttributeService(repo, log),
		Option:        NewOptionService(repo, log),
	}
}

type Authorization interface {
	CreateToken(user models.User, tokenType string, expiresAt time.Time) (*models.Token, error)
	GenerateTokens(user models.User) (*models.Token, *models.Token, error)
	ParseToken(token string) (*jwtCustomClaim, error)
	Login(input models.LoginRequest) (*models.Token, *models.Token, error)
	SignUp(input models.SignUpRequest) (*models.Token, *models.Token, error)
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

type Category interface {
	Create(category models.CategoryCreateRequest) (int64, error)
	GetList(pagination *models.Pagination) ([]models.Category, error)
	GetById(id int64) (models.Category, error)
	Update(category models.CategoryUpdateRequest) error
	DeleteById(id int64) error
}

type Product interface {
	Create(product models.ProductCreateRequest) (int64, error)
	GetList(pagination *models.Pagination) ([]models.Product, error)
	GetById(id int64) (models.Product, error)
	Update(product models.ProductUpdateRequest) error
	DeleteById(id int64) error
	AddAttributeToProduct(productId, attributeId int64) error
	RemoveAttributeFromProduct(productId, attributeId int64) error
	AddToCart(userId int64, request models.CartProductCreateRequest) error
}

type Attribute interface {
	Create(attribute models.AttributeCreateRequest) (int64, error)
	GetList(pagination *models.Pagination) ([]models.Attribute, error)
	GetById(id int64) (models.Attribute, error)
	Update(attribute models.AttributeUpdateRequest) error
	DeleteById(id int64) error
}

type Option interface {
	Create(option models.OptionCreateRequest) (int64, error)
	GetList(pagination *models.Pagination, filters map[string]interface{}) ([]models.Option, error)
	Get(attributeId, optionId int64) (models.Option, error)
	Update(option models.OptionUpdateRequest) error
	Delete(attributeId, optionId int64) error
}
