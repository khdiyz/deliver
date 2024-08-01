package repository

import (
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/pkg/logger"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type OptionRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewOptionRepo(db *sqlx.DB, log logger.Logger) *OptionRepo {
	return &OptionRepo{
		db:  db,
		log: log,
	}
}

func (r *OptionRepo) Create(option models.OptionCreateRequest) (int64, error) {
	var id int64

	query := "INSERT INTO options (name, attribute_id) VALUES ($1, $2) RETURNING id;"

	if err := r.db.QueryRow(query, option.Name, option.AttributeId).Scan(&id); err != nil {
		r.log.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *OptionRepo) GetList(pagination *models.Pagination, filters map[string]interface{}) ([]models.Option, error) {
	var (
		options       []models.Option
		err           error
		filterClauses []string
		args          []interface{}
		counter       int
	)

	countQuery := "SELECT count(id) FROM options WHERE TRUE "

	query := `
	SELECT
		id,
		name,
		attribute_id
	FROM options
	WHERE TRUE `

	if attributeId, ok := filters["attribute-id"]; ok {
		counter++
		filterClauses = append(filterClauses, "attribute_id = $"+strconv.Itoa(counter))
		args = append(args, attributeId)
	}

	if len(filterClauses) > 0 {
		countQuery += " AND " + strings.Join(filterClauses, " AND ")
		query += " AND " + strings.Join(filterClauses, " AND ")
	}

	query += " LIMIT $" + strconv.Itoa(counter+1) + " OFFSET $" + strconv.Itoa(counter+2)

	args = append(args, pagination.Limit, pagination.Offset)

	err = getListCount(r.db, &r.log, pagination, countQuery, args[:len(args)-2])
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var option models.Option
		if rows.Scan(
			&option.Id,
			&option.Name,
			&option.AttributeId,
		); err != nil {
			r.log.Error(err)
			return nil, err
		}

		options = append(options, option)
	}

	return options, nil
}

func (r *OptionRepo) GetById(id int64) (models.Option, error) {
	var option models.Option

	query := `
	SELECT
		id,
		name,
		attribute_id
	FROM options 
	WHERE
		id = $1;`

	err := r.db.QueryRow(query, id).Scan(&option.Id, &option.Name, &option.AttributeId)
	if err != nil {
		r.log.Error(err)
		return models.Option{}, err
	}

	return option, nil
}

func (r *OptionRepo) Update(option models.OptionUpdateRequest) error {
	query := `
	UPDATE options
	SET
		name = $2,
		attribute_id = $3
	WHERE
		id = $1;`

	row, err := r.db.Exec(query, option.Id, option.Name, option.AttributeId)
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

func (r *OptionRepo) DeleteById(id int64) error {
	query := "DELETE FROM options WHERE id = $1;"

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
