package handler

import (
	"log"
	"net/http"

	"github.com/itsgitz/grpc-lab/go/services/common/genproto/orders"
	"github.com/itsgitz/grpc-lab/go/services/common/util"
	"github.com/itsgitz/grpc-lab/go/services/orders/types"
)

type OrdersHTTPHandler struct {
	orderService types.OrderService
}

func NewHTTPOrdersHandler(orderService types.OrderService) *OrdersHTTPHandler {
	return &OrdersHTTPHandler{
		orderService: orderService,
	}
}

func (h *OrdersHTTPHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHTTPHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Println("Request", req.CustomerID)

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}
	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	util.WriteJSON(w, http.StatusOK, res)
}
