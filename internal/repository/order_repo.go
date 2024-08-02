package repository

import (
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/pkg/logger"
	"encoding/json"
	"strconv"
	"strings"

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

func (r *OrderRepo) GetById(id int64) (models.Order, error) {
	var order models.Order

	query := `
	SELECT
	    o.id,
	    COALESCE(o.courier_id, 0),
	    o.reciever_id,
	    o.location_x,
	    o.location_y,
	    o.address,
	    o.status,
	    o.ordered_at,
	    COALESCE(
	        json_agg(
	            json_build_object(
	                'product_id', op.product_id,
					'product_name', p.name,
	                'attributes', op.product_attributes,
	                'quantity', op.quantity
	            )
	        ) FILTER (WHERE op.id IS NOT NULL), '[]'::json
	    ) AS products
	FROM orders o
	JOIN order_products op ON o.id = op.order_id
	JOIN products p ON op.product_id = p.id
	WHERE 
		o.id = $1
	GROUP BY o.id, o.courier_id, o.reciever_id, o.location_x, o.location_y, o.address, o.status, o.ordered_at;`

	var productsData []byte
	if err := r.db.QueryRow(query, id).Scan(
		&order.Id,
		&order.CourierId,
		&order.RecieverId,
		&order.LocationX,
		&order.LocationY,
		&order.Address,
		&order.Status,
		&order.OrderedAt,
		&productsData,
	); err != nil {
		r.log.Error(err)
		return models.Order{}, err
	}

	err := json.Unmarshal(productsData, &order.Products)
	if err != nil {
		r.log.Error(err)
		return models.Order{}, err
	}

	return order, nil
}

func (r *OrderRepo) GetList(pagination *models.Pagination, filters map[string]interface{}) ([]models.Order, error) {
	var (
		orders        []models.Order
		err           error
		filterClauses []string
		args          []interface{}
		counter       int
	)

	countQuery := "SELECT count(o.id) FROM orders o JOIN users u ON o.reciever_id = u.id WHERE true "

	query := `
	SELECT
		o.id,
		COALESCE(o.courier_id, 0),
		o.reciever_id,
		u.full_name,
		o.location_x,
		o.location_y,
		o.address,
		o.status,
		o.ordered_at
	FROM orders o
	JOIN users u ON o.reciever_id = u.id
	WHERE true `

	if orderStatus, ok := filters["status"]; ok {
		counter++
		filterClauses = append(filterClauses, "o.status = $"+strconv.Itoa(counter))
		args = append(args, orderStatus)
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
		var order models.Order
		if rows.Scan(
			&order.Id,
			&order.CourierId,
			&order.RecieverId,
			&order.RecieverName,
			&order.LocationX,
			&order.LocationY,
			&order.Address,
			&order.Status,
			&order.OrderedAt,
		); err != nil {
			r.log.Error(err)
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepo) UpdateById(order models.OrderUpdateRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.log.Error(err)
		return err
	}

	updateOrderQuery := `
	UPDATE orders
	SET
		reciever_id = $2,
		courier_id = $3,
		location_x = $4,
		location_y = $5,
		address = $6,
		status = $7
	WHERE 
		id = $1;`

	if _, err = tx.Exec(updateOrderQuery,
		order.Id,
		order.RecieverId,
		order.CourierId,
		order.LocationX,
		order.LocationY,
		order.Address,
		order.Status,
	); err != nil {
		r.log.Error(err)
		tx.Rollback()
		return err
	}

	updateOrderProductQuery := `
	UPDATE order_products
	SET	
		product_id = $2,
		product_attributes = $3,
		quantity = $4
	WHERE 
		order_id = $1;`

	for _, product := range order.Products {
		productAttributes, err := json.Marshal(product.Attributes)
		if err != nil {
			r.log.Error(err)
			tx.Rollback()
			return err
		}

		if _, err = tx.Exec(updateOrderProductQuery,
			order.Id,
			product.ProductId,
			productAttributes,
			product.Quantity,
		); err != nil {
			r.log.Error(err)
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
