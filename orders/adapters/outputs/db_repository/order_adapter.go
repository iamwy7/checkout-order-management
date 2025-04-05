package db_repository

import (
	"database/sql"

	"github.com/iamwy7/meli-challenge/orders/application/domain"
	ports "github.com/iamwy7/meli-challenge/orders/application/ports/outputs/db_repository"
)

type OrderAdapter struct {
	db *sql.DB
}

func NewMySqlOrderAdapterFactory(db *sql.DB) (ports.OrderRepository, error) {
	return &OrderAdapter{db: db}, nil
}

func (oa *OrderAdapter) CreateOrder(order domain.Order) (*domain.OrderInterface, error) {
	//TODO
	return nil, nil
}

func (oa *OrderAdapter) GetOrderById(orderId string) (*domain.OrderInterface, error) {
	//TODO
	return nil, nil
}

func (oa *OrderAdapter) GetProductsByOrderId(orderId string) (*[]domain.ProductInterface, error) {
	//TODO
	return nil, nil
}

func (oa *OrderAdapter) GetDistributionCentersByProductId(productId string) (*[]domain.DistributionCenterInterface, error) {
	//TODO
	return nil, nil
}
