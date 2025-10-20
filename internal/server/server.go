package server

import (
	"context"
	"my-service/internal/service"
	"my-service/internal/storage"
	pb "my-service/pkg/api/test"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	Service *service.OrderService
}

func NewServer() *Server {
	// Создаем хранилище (in-memory)
	storage := storage.NewMemoryStorage()
	
	// Создаем сервисный слой с зависимостью от хранилища
	service := service.NewOrderService(storage)
	
	// Создаем и возвращаем gRPC сервер
	return &Server{
		Service: service,
	}
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return s.Service.CreateOrder(ctx, req)
}

func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return s.Service.GetOrder(ctx, req)
}

func (s *Server) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	return s.Service.UpdateOrder(ctx, req)
}

func (s *Server) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	return s.Service.DeleteOrder(ctx, req)
}

func (s *Server) ListOrder(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	return s.Service.ListOrder(ctx, req)
}