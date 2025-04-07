package usecase

import (
	out_ports "github.com/iamwy7/meli-challenge/orders/application/ports/outputs"
	"github.com/iamwy7/meli-challenge/orders/application/usecase/dtos"
)

type GetOrderUseCase struct {
	order_repo out_ports.OrderRepository
}

func NewGetOrderUseCase(order_repo out_ports.OrderRepository) *GetOrderUseCase {
	return &GetOrderUseCase{
		order_repo: order_repo}
}

func (uc *GetOrderUseCase) Execute(orderId string) (*dtos.CreatedOrderOutputDTO, error) {
	order, err := uc.order_repo.GetAggregatedOrderById(orderId)
	if err != nil {
		return nil, err
	}
	return MapOrderToOrderOutputDto(order), nil
}
