package out_ports

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
)

type OrderRepository interface {
	CreateAggregatedOrder(order domain.Order) error
	CreateOrder(order domain.Order) error
	CreateProduct(product domain.Product) error
	CreateDistributuionCenter(dc domain.DistributionCenter) error
	LinkProductToOrder(orderId string, prodId string, quantityOrdered int) error
	LinkDistributionCenterToProduct(prodId string, dcId string) error
	GetAggregatedOrderById(orderId string) (*domain.Order, error)
	GetProductsByOrderId(orderId string) (*[]domain.Product, error)
	GetOrderById(orderId string) (*domain.Order, error)
}
