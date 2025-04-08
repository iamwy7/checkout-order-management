package mocks

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
	"github.com/stretchr/testify/mock"
)

type MockDistributionCentersRepository struct {
	mock.Mock
}

func (m *MockDistributionCentersRepository) GetDCsByItemId(itemId string) (*[]domain.DistributionCenter, error) {
	args := m.Called(itemId)
	return args.Get(0).(*[]domain.DistributionCenter), args.Error(1)
}
