package ports

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
)

type OrderRepository interface {
	CreateOrder(order domain.Order) (*domain.OrderInterface, error)
	GetOrderById(orderId string) (*domain.OrderInterface, error)
	GetProductsByOrderId(orderId string) (*[]domain.ProductInterface, error)
	GetDistributionCentersByProductId(productId string) (*[]domain.DistributionCenterInterface, error)
}
