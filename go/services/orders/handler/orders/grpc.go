package handler

import (
	"context"

	"github.com/itsgitz/grpc-lab/go/services/common/genproto/orders"
	"github.com/itsgitz/grpc-lab/go/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGRPCHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGRPCOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	grpcHandler := &OrdersGRPCHandler{
		ordersService: ordersService,
	}
	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrdersGRPCHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}
