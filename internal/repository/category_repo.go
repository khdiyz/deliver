package repository

import (
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/pkg/logger"
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type CategoryRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewCategoryRepo(db *sqlx.DB, log logger.Logger) *CategoryRepo {
	return &CategoryRepo{
		db:  db,
		log: log,
	}
}

func (r *CategoryRepo) Create(category models.CategoryCreateRequest) (int64, error) {
	var id int64

	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id;"

	if err := r.db.QueryRow(query, category.Name).Scan(&id); err != nil {
		r.log.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *CategoryRepo) GetList(pagination *models.Pagination) ([]models.Category, error) {
	var (
		categories []models.Category
		err        error
	)

	countQuery := "SELECT count(id) FROM categories;"
	err = getListCount(r.db, &r.log, pagination, countQuery, nil)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	query := `
	SELECT
	    c.id,
	    c.name,
	    COALESCE(
	        json_agg(
	            json_build_object(
	                'id', p.id,
	                'name', p.name,
	                'description', p.description, 
	                'photo', p.photo,
	                'price', p.price
	            ) 
	        ) FILTER (WHERE p.id IS NOT NULL), '[]'::json
	    ) AS products
	FROM categories c
	LEFT JOIN products p ON c.id = p.category_id 
	GROUP BY c.id, c.name
	LIMIT $1 OFFSET $2;`

	rows, err := r.db.Query(query, pagination.Limit, pagination.Offset)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var (
			category     models.Category
			productsData []byte
		)
		if rows.Scan(
			&category.Id,
			&category.Name,
			&productsData,
		); err != nil {
			r.log.Error(err)
			return nil, err
		}

		err = json.Unmarshal(productsData, &category.Products)
		if err != nil {
			r.log.Error(err)
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepo) GetById(id int64) (models.Category, error) {
	var category models.Category

	query := `
	SELECT
		id,
		name
	FROM categories 
	WHERE
		id = $1;`

	err := r.db.QueryRow(query, id).Scan(&category.Id, &category.Name)
	if err != nil {
		r.log.Error(err)
		return models.Category{}, err
	}

	return category, nil
}

func (r *CategoryRepo) Update(category models.CategoryUpdateRequest) error {
	query := `
	UPDATE categories
	SET
		name = $2
	WHERE
		id = $1;`

	row, err := r.db.Exec(query, category.Id, category.Name)
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

func (r *CategoryRepo) DeleteById(id int64) error {
	query := "DELETE FROM categories WHERE id = $1;"

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
