package usecase

import (
	out_ports "github.com/iamwy7/meli-challenge/orders/application/ports/outputs"
	"github.com/iamwy7/meli-challenge/orders/application/usecase/dtos"
)

type GetOrderUseCase struct {
	orderRepo out_ports.OrderRepository
}

func NewGetOrderUseCase(orderRepo out_ports.OrderRepository) *GetOrderUseCase {
	return &GetOrderUseCase{
		orderRepo: orderRepo}
}

func (uc *GetOrderUseCase) Execute(orderId string) (*dtos.CreatedOrderOutputDTO, error) {
	order, err := uc.orderRepo.GetAggregatedOrderById(orderId)
	if err != nil {
		return nil, err
	}
	return MapOrderToOrderOutputDto(order), nil
}
