package types

import (
	"context"

	"github.com/itsgitz/grpc-lab/go/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
