package repository

import (
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/pkg/logger"
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewOrderRepo(db *sqlx.DB, log logger.Logger) *OrderRepo {
	return &OrderRepo{
		db:  db,
		log: log,
	}
}

func (r *OrderRepo) Create(order models.OrderCreateRequest) (int64, error) {
	var orderId int64

	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error(err)
		return 0, err
	}

	createOrderQuery := `
	INSERT INTO orders(
		reciever_id,
		location_x,
		location_y,
		address,
		status
	) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	if err = tx.QueryRow(createOrderQuery,
		order.RecieverId,
		order.LocationX,
		order.LocationY,
		order.Address,
		constants.OrderStatusPickedUp,
	).Scan(&orderId); err != nil {
		r.log.Error(err)
		tx.Rollback()
		return 0, err
	}

	createOrderProductQuery := `
	INSERT INTO order_products(
		order_id,	
		product_id,
		product_attributes,
		quantity
	) VALUES ($1, $2, $3, $4);`

	for _, product := range order.Products {
		productAttributes, err := json.Marshal(product.Attributes)
		if err != nil {
			r.log.Error(err)
			tx.Rollback()
			return 0, err
		}

		if _, err = tx.Exec(createOrderProductQuery,
			orderId,
			product.ProductId,
			productAttributes,
			product.Quantity,
		); err != nil {
			r.log.Error(err)
			tx.Rollback()
			return 0, err
		}
	}

	return orderId, tx.Commit()
}
