package main

import (
	"log"
	"net/http"

	handler "github.com/itsgitz/grpc-lab/go/services/orders/handler/orders"
	services "github.com/itsgitz/grpc-lab/go/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHTTPServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := services.NewOrderService()
	orderHandler := handler.NewHTTPOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting HTTP server on port", s.addr)

	return http.ListenAndServe(s.addr, router)
}
