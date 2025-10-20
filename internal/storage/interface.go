package storage

import "context"

type Storage interface {
	CreateOrder(ctx context.Context, item string, quantity int32) (string, error)
	GetOrder(ctx context.Context, id string) (*Order, error)
	UpdateOrder(ctx context.Context, id string, item string, quantity int32) (*Order, error)
	DeleteOrder(ctx context.Context, id string) (bool, error)
	ListOrder(ctx context.Context) ([]*Order, error)
}