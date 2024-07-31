package repository

import (
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/pkg/logger"
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewProductRepo(db *sqlx.DB, log logger.Logger) *ProductRepo {
	return &ProductRepo{
		db:  db,
		log: log,
	}
}

func (r *ProductRepo) Create(product models.ProductCreateRequest) (int64, error) {
	var id int64

	query := `
	INSERT INTO products (
		name,
		category_id,
		description,
		photo,
		price 
	) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	if err := r.db.QueryRow(query,
		product.Name,
		product.CategoryId,
		product.Description,
		product.Photo,
		product.Price,
	).Scan(&id); err != nil {
		r.log.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *ProductRepo) GetList(pagination *models.Pagination) ([]models.Product, error) {
	var (
		products []models.Product
		err      error
	)

	countQuery := "SELECT count(id) FROM products;"
	err = getListCount(r.db, &r.log, pagination, countQuery, nil)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	query := `
	SELECT
		id,
		name,
		category_id,
		description,
		photo,
		price
	FROM products 
	LIMIT $1 OFFSET $2;`

	rows, err := r.db.Query(query, pagination.Limit, pagination.Offset)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	for rows.Next() {
		var product models.Product
		if rows.Scan(
			&product.Id,
			&product.Name,
			&product.CategoryId,
			&product.Description,
			&product.Photo,
			&product.Price,
		); err != nil {
			r.log.Error(err)
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepo) GetById(id int64) (models.Product, error) {
	var product models.Product

	query := `
	SELECT
	    p.id,
	    p.name,
	    p.category_id,
	    p.description,
	    p.photo,
	    p.price,
	    COALESCE(
	        json_agg(
	            json_build_object(
	                'id', a.id,
	                'name', a.name,
	                'options', (
	                    SELECT COALESCE(json_agg(
	                        json_build_object(
	                            'id', o.id,
	                            'name', o.name
	                        )
	                    ), '[]'::json)
	                    FROM options o
	                    WHERE o.attribute_id = a.id
	                )
	            )
	        ) FILTER (WHERE a.id IS NOT NULL), '[]'::json
	    ) AS attributes
	FROM products p
	LEFT JOIN product_attributes pa ON p.id = pa.product_id
	LEFT JOIN attributes a ON pa.attribute_id = a.id
	WHERE p.id = $1
	GROUP BY p.id, p.name, p.category_id, p.description, p.photo, p.price;`

	var attributesData []byte

	if err := r.db.QueryRow(query, id).Scan(
		&product.Id,
		&product.Name,
		&product.CategoryId,
		&product.Description,
		&product.Photo,
		&product.Price,
		&attributesData,
	); err != nil {
		r.log.Error(err)
		return models.Product{}, err
	}

	err := json.Unmarshal(attributesData, &product.Attributes)
	if err != nil {
		r.log.Error(err)
		return models.Product{}, err
	}

	return product, nil
}

func (r *ProductRepo) Update(product models.ProductUpdateRequest) error {
	query := `
	UPDATE products
	SET
		name = $2,
		category_id = $3,
		description = $4,
		photo = $5,
		price = $6
	WHERE
		id = $1;`

	row, err := r.db.Exec(query,
		product.Id,
		product.Name,
		product.CategoryId,
		product.Description,
		product.Photo,
		product.Price,
	)
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
		return constants.ErrorNoRowsAffected
	}

	return nil
}

func (r *ProductRepo) DeleteById(id int64) error {
	query := "DELETE FROM products WHERE id = $1;"

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
		return constants.ErrorNoRowsAffected
	}

	return nil
}
