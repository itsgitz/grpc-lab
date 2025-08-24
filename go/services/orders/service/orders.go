package services

import (
	"context"

	"github.com/itsgitz/grpc-lab/go/services/common/genproto/orders"
)

var ordersDB = make([]*orders.Order, 0)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDB = append(ordersDB, order)
	return nil
}
