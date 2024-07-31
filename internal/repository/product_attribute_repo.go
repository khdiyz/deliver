package repository

import (
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type ProductAttributeRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewProductAttributeRepo(db *sqlx.DB, log logger.Logger) *ProductAttributeRepo {
	return &ProductAttributeRepo{
		db:  db,
		log: log,
	}
}

func (r *ProductAttributeRepo) Create(request models.AddAttributeToProduct) (int64, error) {
	var id int64

	query := `
	INSERT INTO product_attributes (
		product_id,
		attribute_id
	) VALUES ($1, $2) RETURNING id;`

	err := r.db.QueryRow(query, request.ProductId, request.AttributeId).Scan(&id)
	if err != nil {
		r.log.Error(err)
		return 0, err
	}

	return id, nil
}

func (r *ProductAttributeRepo) GetByProductIdAndAttributeId(productId, attributeId int64) (models.ProductAttribute, error) {
	var productAttribute models.ProductAttribute

	query := `
	SELECT
		id,
		product_id,
		attribute_id
	FROM product_attributes
	WHERE
		product_id = $1
		AND attribute_id = $2;`

	if err := r.db.QueryRow(query, productId, attributeId).Scan(
		&productAttribute.Id,
		&productAttribute.ProductId,
		&productAttribute.AttributeId,
	); err != nil {
		r.log.Error(err)
		return models.ProductAttribute{}, err
	}

	return productAttribute, nil
}

func (r *ProductAttributeRepo) DeleteByProductIdAndAttributeId(productId, attributeId int64) error {
	query := `
	DELETE FROM product_attributes
 	WHERE
		product_id = $1
		AND attribute_id = $2;`

	row, err := r.db.Exec(query, productId, attributeId)
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
