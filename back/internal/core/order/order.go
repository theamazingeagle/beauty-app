package Order

import (
	"beauty/internal/types"
	"errors"
	"log"
)

type Storage interface {
	GetOrderByID(ID types.OrderID) (types.Order, bool, error)
	GetOrders() ([]types.Order, error)
	CreateOrder(newOrder types.Order) error
	UpdateOrder(Order types.Order) error
	DeleteOrderByID(ID types.OrderID) error
}

var (
	ErrOrderNotFound    = errors.New("Order not found")
	ErrFailedToFetch    = errors.New("Failed to fetch order info")
	ErrFailedToFetchAll = errors.New("Failed to fetch all orders info")
	ErrFailedToCreate   = errors.New("Failed to create new order record")
	ErrFailedToUpdate   = errors.New("Failed to update order info")
	ErrFailedToDelete   = errors.New("Failed to delete order info")
)

type OrderModel struct {
	storage Storage
}

func New(s Storage) OrderModel {
	return OrderModel{storage: s}
}

func (om *OrderModel) Get(ID types.OrderID) (types.Order, error) {
	Order, exist, err := om.storage.GetOrderByID(ID)
	if err != nil {
		log.Println("Failed to get order info")
		return types.Order{}, ErrFailedToFetch
	}
	if !exist {
		log.Println("Order not found")
		return types.Order{}, ErrOrderNotFound
	}
	return Order, nil
}

func (om *OrderModel) GetAll() ([]types.Order, error) {
	allOrders, err := om.storage.GetOrders()
	if err != nil {
		log.Println("Failed to get all orders info")
		return nil, ErrFailedToFetchAll
	}
	return allOrders, nil
}

func (om *OrderModel) Create(Order types.Order) error {
	err := om.storage.CreateOrder(Order)
	if err != nil {
		log.Println("Failed to create order record")
		return ErrFailedToCreate
	}
	return nil
}

func (om *OrderModel) Update(Order types.Order) error {
	err := om.storage.UpdateOrder(Order)
	if err != nil {
		log.Println("Failed to update order record")
		return ErrFailedToUpdate
	}
	return nil
}

func (om *OrderModel) Delete(ID types.OrderID) error {
	err := om.storage.DeleteOrderByID(ID)
	if err != nil {
		log.Println("Failed to delete order record")
		return ErrFailedToDelete
	}
	return nil
}
