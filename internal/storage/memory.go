package storage

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
)

const ErrOrderNotFound = "order not found"

type Order struct {
	ID 	     string
	Item   	 string
	Quantity int32
}

type MemoryStorage struct {
	mu sync.RWMutex
	orders map[string]*Order
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		orders: make(map[string]*Order),
	}
}

func (s *MemoryStorage) CreateOrder(ctx context.Context, item string, quantity int32) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()

	order := &Order{
		ID: id,
		Item: item,
		Quantity: quantity,
	}

	s.orders[id] = order
	
	return id, nil
}

func (s *MemoryStorage) GetOrder(ctx context.Context, id string) (*Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	order, exists := s.orders[id]
	if !exists {
		return nil, errors.New(ErrOrderNotFound)
	}

	return order, nil
}

func (s *MemoryStorage) UpdateOrder(ctx context.Context, id string, item string, quantity int32) (*Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, exists := s.orders[id]
	if !exists {
		return nil, errors.New(ErrOrderNotFound)
	}

	order.Item = item
	order.Quantity = quantity

	return order, nil
}

func (s *MemoryStorage) DeleteOrder(ctx context.Context, id string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.orders[id]
	if !exists {
		return false, errors.New(ErrOrderNotFound)
	}

	delete(s.orders, id)

	return true, nil
}

func (s *MemoryStorage) ListOrder(ctx context.Context) ([]*Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var Orders []*Order
	for _, order := range s.orders {
		Orders = append(Orders, order)
	}

	return Orders, nil
}