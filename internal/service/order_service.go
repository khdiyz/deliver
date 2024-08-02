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
