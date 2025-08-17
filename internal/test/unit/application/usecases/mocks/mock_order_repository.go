package mocks

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateAggregatedOrder(order domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}
func (m *MockOrderRepository) CreateOrder(order domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) CreateProduct(product domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockOrderRepository) CreateDistributuionCenter(dc domain.DistributionCenter) error {
	args := m.Called(dc)
	return args.Error(0)
}

func (m *MockOrderRepository) LinkProductToOrder(orderId string, prodId string, quantityOrdered int) error {
	args := m.Called(orderId, prodId, quantityOrdered)
	return args.Error(0)
}

func (m *MockOrderRepository) LinkDistributionCenterToProduct(prodId string, dcId string) error {
	args := m.Called(prodId, dcId)
	return args.Error(0)
}

func (m *MockOrderRepository) GetAggregatedOrderById(orderId string) (*domain.Order, error) {
	args := m.Called(orderId)
	return args.Get(0).(*domain.Order), args.Error(1)
}

func (m *MockOrderRepository) GetProductsByOrderId(orderId string) (*[]domain.Product, error) {
	args := m.Called(orderId)
	return args.Get(0).(*[]domain.Product), args.Error(1)
}

func (m *MockOrderRepository) GetOrderById(orderId string) (*domain.Order, error) {
	args := m.Called(orderId)
	return args.Get(0).(*domain.Order), args.Error(1)
}
