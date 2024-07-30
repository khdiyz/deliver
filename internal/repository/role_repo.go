package repository

import (
	"deliver/models"
	"deliver/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type RoleRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewRoleRepo(db *sqlx.DB, log logger.Logger) *RoleRepo {
	return &RoleRepo{
		db:  db,
		log: log,
	}
}

func (r *RoleRepo) GetList(pagination *models.Pagination) ([]models.Role, error) {
	var (
		roles []models.Role
		err   error
	)

	countQuery := "SELECT count(id) FROM roles;"
	err = getListCount(r.db, &r.log, pagination, countQuery, nil)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description
	FROM roles 
	LIMIT $1 OFFSET $2;`

	rows, err := r.db.Query(query, pagination.Limit, pagination.Offset)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var role models.Role
		if rows.Scan(
			&role.Id,
			&role.Name,
			&role.Description,
		); err != nil {
			r.log.Error(err)
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (r *RoleRepo) GetById(id int64) (models.Role, error) {
	var role models.Role

	query := `
	SELECT
		id,
		name,
		COALESCE(description, '') AS description
	FROM roles
	WHERE
		id = $1;`

	if err := r.db.QueryRow(query, id).Scan(&role.Id, &role.Name, &role.Description); err != nil {
		return models.Role{}, err
	}

	return role, nil
}
