package service

import (
	"context"
	"my-service/internal/storage"
	pb "my-service/pkg/api/test"
)

type OrderService struct {
	Storage storage.Storage
}

func NewOrderService(storage storage.Storage) *OrderService {
	return &OrderService{Storage: storage}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	id, err := s.Storage.CreateOrder(ctx, req.Item, req.Quantity)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{Id: id}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	order, err := s.Storage.GetOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetOrderResponse{
		Order: &pb.Order{
			Id: order.ID,
			Item: order.Item,
			Quantity: order.Quantity,
		},
	}, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	order, err := s.Storage.UpdateOrder(ctx, req.Id, req.Item, req.Quantity)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateOrderResponse{
		Order: &pb.Order{
			Id: order.ID,
			Item: order.Item,
			Quantity: order.Quantity,
		},
	}, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	success, err := s.Storage.DeleteOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteOrderResponse{Success: success}, nil
}

func (s *OrderService) ListOrder(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.Storage.ListOrder(ctx)
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id: order.ID,
			Item: order.Item,
			Quantity: order.Quantity,
		})
	}

	return &pb.ListOrdersResponse{Orders: pbOrders}, nil
} 