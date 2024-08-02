package service

import (
	"deliver/internal/constants"
	"deliver/internal/models"
	"deliver/internal/repository"
	"deliver/internal/ws"
	"deliver/pkg/logger"
	"errors"
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
	for i := range input.Products {
		if input.Products[i].Quantity < 1 {
			return 0, serviceError(errors.New("quantity must be greater 1"), codes.InvalidArgument)
		}
	}

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

func (s *OrderService) ReceiveOrderCourier(orderId, courierId int64) error {
	order, err := s.repo.Order.GetById(orderId)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	if order.Status != constants.OrderStatusPickedUp {
		return serviceError(errors.New("already recieved"), codes.InvalidArgument)
	}

	s.repo.Order.UpdateById(models.OrderUpdateRequest{
		Id:         orderId,
		RecieverId: order.RecieverId,
		CourierId:  courierId,
		LocationX:  order.LocationX,
		LocationY:  order.LocationY,
		Address:    order.Address,
		Status:     constants.OrderStatusOnDelivery,
	})

	update := ws.OrderStatusUpdate{
		OrderID: fmt.Sprintf("%d", orderId),
		Status:  constants.OrderStatusOnDelivery,
	}

	ws.BroadcastOrderStatusUpdate(update)

	return nil
}

func (s *OrderService) FinishOrderCourier(orderId, courierId int64) error {
	order, err := s.repo.Order.GetById(orderId)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	if order.Status != constants.OrderStatusOnDelivery {
		return serviceError(errors.New("already delivered"), codes.InvalidArgument)
	}

	s.repo.Order.UpdateById(models.OrderUpdateRequest{
		Id:         orderId,
		RecieverId: order.RecieverId,
		CourierId:  courierId,
		LocationX:  order.LocationX,
		LocationY:  order.LocationY,
		Address:    order.Address,
		Status:     constants.OrderStatusDelivered,
	})

	update := ws.OrderStatusUpdate{
		OrderID: fmt.Sprintf("%d", orderId),
		Status:  constants.OrderStatusDelivered,
	}

	ws.BroadcastOrderStatusUpdate(update)

	return nil
}

func (s *OrderService) PaymentCollectedOrderCourier(orderId, courierId int64) error {
	order, err := s.repo.Order.GetById(orderId)
	if err != nil {
		return serviceError(err, codes.Internal)
	}

	if order.Status != constants.OrderStatusDelivered {
		return serviceError(errors.New("already completed"), codes.InvalidArgument)
	}

	s.repo.Order.UpdateById(models.OrderUpdateRequest{
		Id:         orderId,
		RecieverId: order.RecieverId,
		CourierId:  courierId,
		LocationX:  order.LocationX,
		LocationY:  order.LocationY,
		Address:    order.Address,
		Status:     constants.OrderStatusPaymentCollected,
	})

	update := ws.OrderStatusUpdate{
		OrderID: fmt.Sprintf("%d", orderId),
		Status:  constants.OrderStatusPaymentCollected,
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
