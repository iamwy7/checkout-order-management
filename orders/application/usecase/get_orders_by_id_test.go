package usecase

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/iamwy7/meli-challenge/orders/application/domain"
	"github.com/iamwy7/meli-challenge/orders/application/shared"
	"github.com/iamwy7/meli-challenge/orders/application/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func createAnOrder() *domain.Order {
	distributionCenter := &domain.DistributionCenter{
		Id:              "e4f5a7d2-6c9f-1d8b-0c2a-3d4e5b2a7f6c",
		Name:            "CD15",
		Zone:            shared.North2,
		State:           "SP",
		Status:          shared.ACTIVE,
		ProductQuantity: 6,
	}

	product := &domain.Product{
		Id:                 uuid.New().String(),
		CatalogProductId:   "0cb0a0d0-a815-41c6-9801-d4704c381cee",
		Name:               "Product Name",
		Price:              27.77,
		Quantity:           2,
		DistributionCenter: *distributionCenter,
	}
	var products []domain.Product
	products = append(products, *product)

	return &domain.Order{
		Id:            "8c8f5d8f-bcec-478c-8d56-5c8390d0f938",
		Zone:          shared.South1,
		State:         "SP",
		Status:        shared.ACTIVE,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Products:      products,
		ProductsCount: len(products),
	}
}

func TestGetOrdersUseCase(t *testing.T) {
	// Arrange
	mockOrderRepo := new(mocks.MockOrderRepository)
	getorderUseCase := NewGetOrderUseCase(mockOrderRepo)
	validOrderId := "8c8f5d8f-bcec-478c-8d56-5c8390d0f938"

	// Act
	mockOrderRepo.On("GetAggregatedOrderById", validOrderId).Return(createAnOrder(), nil)
	output, err := getorderUseCase.Execute(validOrderId)

	// Assert
	assert.NotNil(t, output)
	assert.Nil(t, err)
	assert.Equal(t, validOrderId, output.Id)

	// Assert that the mock repository was called
	mockOrderRepo.AssertExpectations(t)
}
