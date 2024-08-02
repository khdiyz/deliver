package service

import (
	"deliver/internal/models"
	"deliver/internal/repository"
	"deliver/internal/ws"
	"deliver/pkg/logger"
	"fmt"

	"google.golang.org/grpc/codes"
)

type OrderService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewOrderService(repo repository.Repository, log logger.Logger) *OrderService {
	return &OrderService{
		repo: repo,
		log:  log,
	}
}

func (s *OrderService) CreateOrder(input models.OrderCreateRequest) (int64, error) {
	orderId, err := s.repo.Order.Create(input)
	if err != nil {
		return 0, serviceError(err, codes.Internal)
	}

	update := ws.OrderStatusUpdate{
		OrderID: fmt.Sprintf("%d", orderId),
		Status:  "picked_up",
	}

	ws.BroadcastOrderStatusUpdate(update)

	return orderId, nil
}

func (s *OrderService) ReceiveOrderCourier(orderId int64, input models.OrderCourierRequest) error {
	// Logic to handle receiving the order courier

	// Send message to WebSocket with order_id and status "on_delivery"
	update := ws.OrderStatusUpdate{
		OrderID: fmt.Sprintf("%d", orderId),
		Status:  "on_delivery",
	}

	ws.BroadcastOrderStatusUpdate(update)

	return nil
}

func (s *OrderService) GetById(id int64) (models.Order, error) {
	order, err := s.repo.Order.GetById(id)
	if err != nil {
		return models.Order{}, serviceError(err, codes.Internal)
	}

	return order, nil
}

func (s *OrderService) GetList(pagination *models.Pagination, filters map[string]interface{}) ([]models.Order, error) {
	orders, err := s.repo.Order.GetList(pagination, filters)
	if err != nil {
		return nil, serviceError(err, codes.Internal)
	}

	return orders, nil
}

// func (s *OrderService) UpdateById(order models.OrderUpdateRequest) error {
// 	err := s.repo.Order.UpdateById(order)
// 	if err != nil {
// 		return serviceError(err, codes.Internal)
// 	}

// 	update := ws.OrderStatusUpdate{
// 		OrderID: fmt.Sprintf("%d", order.Id),
// 		Status:  ,
// 	}

// 	ws.BroadcastOrderStatusUpdate(update)

// 	return nil
// }
