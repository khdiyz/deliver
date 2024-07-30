package repository

import (
	"deliver/models"
	"deliver/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewUserRepo(db *sqlx.DB, log logger.Logger) *UserRepo {
	return &UserRepo{
		db:  db,
		log: log,
	}
}

func (r *UserRepo) Create(input models.UserCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO users (
		full_name,
		email,
		password,
		role_id
	) VALUES ($1, $2, $3, $4) RETURNING id;`

	if err := r.db.QueryRow(query,
		input.FullName,
		input.Email,
		input.Password,
		input.RoleId,
	).Scan(&id); err != nil {
		r.log.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) GetByEmail(email string) (models.User, error) {
	var user models.User

	query := `
	SELECT
		u.id,
		u.full_name,
		u.email,
		u.password,
		u.role_id,
		r.name
	FROM users u
	JOIN roles r ON u.role_id = r.id
	WHERE
		u.email = $1;`

	err := r.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.RoleId,
		&user.RoleName,
	)
	if err != nil {
		r.log.Error(err)
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetById(id int64) (models.User, error) {
	var user models.User

	query := `
	SELECT
		u.id,
		u.full_name,
		u.email,
		u.password,
		u.role_id,
		r.name
	FROM users u 
	JOIN roles r ON u.role_id = r.id
	WHERE
		u.id = $1;`

	if err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.RoleId,
		&user.RoleName,
	); err != nil {
		r.log.Error(err)
		return models.User{}, err
	}

	return user, nil
}
