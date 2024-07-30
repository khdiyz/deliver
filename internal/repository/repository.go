package repository

import (
	"deliver/models"
	"deliver/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User
	Role
}

func NewRepository(db *sqlx.DB, log logger.Logger) *Repository {
	return &Repository{
		User: NewUserRepo(db, log),
		Role: NewRoleRepo(db, log),
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
