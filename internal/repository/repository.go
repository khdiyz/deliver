package repository

import (
	"deliver/internal/models"
	"deliver/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User
	Role
	Category
	Product
	Attribute
	Option
	ProductAttribute
}

func NewRepository(db *sqlx.DB, log logger.Logger) *Repository {
	return &Repository{
		User:             NewUserRepo(db, log),
		Role:             NewRoleRepo(db, log),
		Category:         NewCategoryRepo(db, log),
		Product:          NewProductRepo(db, log),
		Attribute:        NewAttributeRepo(db, log),
		Option:           NewOptionRepo(db, log),
		ProductAttribute: NewProductAttributeRepo(db, log),
	}
}

type User interface {
	Create(input models.UserCreateRequest) (int64, error)
	GetByEmail(email string) (models.User, error)
	GetById(id int64) (models.User, error)
}

type Role interface {
	GetList(pagination *models.Pagination) ([]models.Role, error)
	GetById(id int64) (models.Role, error)
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
	GetById(id int64) (models.Option, error)
	Update(option models.OptionUpdateRequest) error
	DeleteById(id int64) error
}

type ProductAttribute interface {
	Create(request models.AddAttributeToProduct) (int64, error)
	GetByProductIdAndAttributeId(productId, attributeId int64) (models.ProductAttribute, error)
	DeleteByProductIdAndAttributeId(productId, attributeId int64) error
}
