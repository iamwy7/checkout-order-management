package out_ports

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
)

type OrderRepository interface {
	CreateOrder(order domain.Order) error
	CreateProduct(product domain.Product) error
	CreateDistributuionCenter(dc domain.DistributionCenter) error
	GetProductsByOrderId(orderId string) (*[]domain.Product, error)
	GetOrderById(orderId string) (*domain.Order, error)
}
