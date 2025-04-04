package ports

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
)

type OrderRepository interface {
	CreateOrder(order domain.Order) (*domain.Order, error)
	GetOrder(orderID string) (*domain.Order, error)
}
