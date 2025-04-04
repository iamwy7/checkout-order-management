package ports

import "github.com/iamwy7/meli-challenge/orders/application/domain"

type OrdersAPI interface {
	CreateOrder(order domain.Order) (*domain.Order, error)
	GetOrderById(orderID string) (*domain.Order, error)
}
