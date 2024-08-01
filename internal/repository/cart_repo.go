package repository

import (
	"deliver/internal/models"
	"deliver/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type CartRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewCartRepo(db *sqlx.DB, log logger.Logger) *CartRepo {
	return &CartRepo{
		db:  db,
		log: log,
	}
}

func (r *CartRepo) Create(userId int64) (int64, error) {
	var cartId int64

	query := "INSERT INTO cart(user_id) VALUES($1) RETURNING id;"

	err := r.db.QueryRow(query, userId).Scan(&cartId)
	if err != nil {
		r.log.Error(err)
		return 0, err
	}

	return cartId, nil
}

func (r *CartRepo) GetCartIdByUserId(userId int64) (int64, error) {
	var cartId int64

	query := "SELECT id FROM cart WHERE user_id = $1;"

	err := r.db.QueryRow(query, userId).Scan(&cartId)
	if err != nil {
		r.log.Error(err)
		return 0, err
	}

	return cartId, err
}

func (r *CartRepo) CreateCartProduct(cartProduct models.CartProductCreateRequest) (int64, error) {
	var cartProductId int64

	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error(err)
		return 0, err
	}

	createCartProductQuery := `
	INSERT INTO cart_products(
		cart_id,
		product_id,
		quantity
	) VALUES ($1, $2, $3) RETURNING id;`

	if err = tx.QueryRow(createCartProductQuery,
		cartProduct.CartId,
		cartProduct.ProductId,
		cartProduct.Quantity,
	).Scan(&cartProductId); err != nil {
		r.log.Error(err)
		tx.Rollback()
		return 0, err
	}

	createCartProductAttributesQuery := `
	INSERT INTO cart_product_attributes (
		cart_product_id,
		attribute_id,
		option_id
	) VALUES ($1, $2, $3);`

	for i := range cartProduct.Attributes {
		if _, err = tx.Exec(createCartProductAttributesQuery,
			cartProductId,
			cartProduct.Attributes[i].AttributeId,
			cartProduct.Attributes[i].OptionId,
		); err != nil {
			r.log.Error(err)
			tx.Rollback()
			return 0, err
		}
	}

	return cartProductId, tx.Commit()
}
