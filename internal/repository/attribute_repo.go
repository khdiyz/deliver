package repository

import (
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/pkg/logger"
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type AttributeRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewAttributeRepo(db *sqlx.DB, log logger.Logger) *AttributeRepo {
	return &AttributeRepo{
		db:  db,
		log: log,
	}
}

func (r *AttributeRepo) Create(attribute models.AttributeCreateRequest) (int64, error) {
	var id int64

	query := "INSERT INTO attributes (name) VALUES ($1) RETURNING id;"

	if err := r.db.QueryRow(query, attribute.Name).Scan(&id); err != nil {
		r.log.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *AttributeRepo) GetList(pagination *models.Pagination) ([]models.Attribute, error) {
	var (
		attributes []models.Attribute
		err        error
	)

	countQuery := "SELECT count(id) FROM attributes;"
	err = getListCount(r.db, &r.log, pagination, countQuery, nil)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		name
	FROM attributes 
	LIMIT $1 OFFSET $2;`

	rows, err := r.db.Query(query, pagination.Limit, pagination.Offset)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var attribute models.Attribute
		if rows.Scan(
			&attribute.Id,
			&attribute.Name,
		); err != nil {
			r.log.Error(err)
			return nil, err
		}

		attributes = append(attributes, attribute)
	}

	return attributes, nil
}

func (r *AttributeRepo) GetById(id int64) (models.Attribute, error) {
	var attribute models.Attribute

	query := `
	SELECT
	    a.id,
	    a.name,
	    COALESCE(json_agg(
	        json_build_object(
	            'id', o.id,
	            'name', o.name
	        )
	    ) FILTER (WHERE o.id IS NOT NULL), '[]') AS options
	FROM attributes a
	LEFT JOIN options o ON a.id = o.attribute_id
	WHERE a.id = $1
	GROUP BY a.id, a.name;`

	var optionsData []byte
	err := r.db.QueryRow(query, id).Scan(&attribute.Id, &attribute.Name, &optionsData)
	if err != nil {
		r.log.Error(err)
		return models.Attribute{}, err
	}

	err = json.Unmarshal(optionsData, &attribute.Options)
	if err != nil {
		r.log.Error(err)
		return models.Attribute{}, err
	}

	return attribute, nil
}

func (r *AttributeRepo) Update(attribute models.AttributeUpdateRequest) error {
	query := `
	UPDATE attributes
	SET
		name = $2
	WHERE
		id = $1;`

	row, err := r.db.Exec(query, attribute.Id, attribute.Name)
	if err != nil {
		r.log.Error(err)
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		r.log.Error(err)
		return err
	}
	if rowAffected == 0 {
		return constants.ErrNoRowsAffected
	}

	return nil
}

func (r *AttributeRepo) DeleteById(id int64) error {
	query := "DELETE FROM attributes WHERE id = $1;"

	row, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Error(err)
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		r.log.Error(err)
		return err
	}
	if rowAffected == 0 {
		return constants.ErrNoRowsAffected
	}

	return nil
}
