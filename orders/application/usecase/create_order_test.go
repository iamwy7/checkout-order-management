package usecase

import (
	"testing"

	"github.com/iamwy7/meli-challenge/orders/application/domain"
	"github.com/iamwy7/meli-challenge/orders/application/shared"
	"github.com/iamwy7/meli-challenge/orders/application/usecase/dtos"
	"github.com/iamwy7/meli-challenge/orders/application/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateEventUseCase(t *testing.T) {
	// Arrange
	mockOrderRepo := new(mocks.MockOrderRepository)
	mockDcRepo := new(mocks.MockDistributionCentersRepository)
	createOrderUseCase := NewCreateOrderUseCase(mockOrderRepo, mockDcRepo)

	var domainDcs = make([]domain.DistributionCenter, 1)
	var dc = &domain.DistributionCenter{
		Id:              "ed7905ca-a47a-46f9-b623-baf06bb7e59e",
		Name:            "CD1",
		Zone:            shared.South1,
		State:           "SP",
		Status:          shared.ACTIVE,
		ProductQuantity: 7,
	}
	domainDcs[0] = *dc

	var productsInputs = make([]dtos.ProductInputDTO, 1)
	productsInput := &dtos.ProductInputDTO{
		Id:       "0f99276b-aa53-44f6-8bb6-5b4ededc9615",
		Name:     "Product Name",
		Price:    27.77,
		Quantity: 2,
	}
	productsInputs[0] = *productsInput
	input := &dtos.CreateOrderInputDTO{
		Products: productsInputs,
		Zone:     "S1",
		State:    "SP",
	}

	mockDcRepo.On("GetDCsByItemId", mock.AnythingOfType("string")).Return(&domainDcs, nil)
	mockOrderRepo.On("CreateAggregatedOrder", mock.AnythingOfType("domain.Order")).Return(nil)
	mockOrderRepo.On("CreateOrder", mock.AnythingOfType("domain.Order")).Return(nil)
	mockOrderRepo.On("CreateProduct", mock.AnythingOfType("domain.Product")).Return(nil)
	mockOrderRepo.On("CreateDistributuionCenter", mock.AnythingOfType("domain.DistributionCenter")).Return(nil)
	mockOrderRepo.On("LinkProductToOrder", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil)
	mockOrderRepo.On("LinkDistributionCenterToProduct", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	output, err := createOrderUseCase.Execute(*input)
	assert.NotNil(t, output)
	assert.Nil(t, err)
}
