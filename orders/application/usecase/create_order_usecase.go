package usecase

import (
	"fmt"

	"github.com/iamwy7/meli-challenge/orders/application/domain"
	out_ports "github.com/iamwy7/meli-challenge/orders/application/ports/outputs"
	usecase "github.com/iamwy7/meli-challenge/orders/application/usecase/dtos"
)

// Precisa criar um DTO que seja o pedido recebido pela api
// Também precisa de um que é o que vai ser retornado pela api

type CreateOrderUseCase struct {
	order_repo out_ports.OrderRepository
	dc_repo    out_ports.DistributionCentersRepository
}

func NewCreateOrderUseCase(order_repo out_ports.OrderRepository, dc_repo out_ports.DistributionCentersRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		order_repo: order_repo,
		dc_repo:    dc_repo,
	}
}
func (uc *CreateOrderUseCase) Execute(input usecase.CreateOrderInputDTO) (*usecase.CreatedOrderOutputDto, error) {
	// Validate Products
	products := make([]domain.Product, len(input.Products))
	for _, product := range input.Products {
		dcs, err := uc.getDCsForEachProduct(product)
		if err != nil {
			return nil, err
		}
		chosenDC, err := uc.getDCWithHighestProductQuantity(dcs)
		if err != nil {
			return nil, err
		}

		if product.Quantity > chosenDC.ProductQuantity {
			return nil, domain.ErrProductQuantityExceeded
		}
		p, err := domain.NewProduct(product.Id, product.Name, product.Price, product.Quantity, *chosenDC)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	// Validate Order
	order, err := domain.NewOrder(products, input.Zone, input.State)
	if err != nil {
		return nil, err
	}
	err = uc.order_repo.CreateAggregatedOrder(*order)
	if err != nil {
		return nil, err
	}

	return MapOrderToOrderOutputDto(order), nil
}

func (uc *CreateOrderUseCase) getDCsForEachProduct(productDto usecase.ProductInputDto) (*[]domain.DistributionCenter, error) {
	dc, err := uc.dc_repo.GetDCsByItemId(productDto.Id)
	if err != nil {
		return nil, err
	}
	if len(*dc) == 0 {
		return nil, domain.ErrProductInvalidId
	}
	return dc, nil
}

func (uc *CreateOrderUseCase) getDCWithHighestProductQuantity(domainDCs *[]domain.DistributionCenter) (*domain.DistributionCenter, error) {
	var maxDC *domain.DistributionCenter
	maxQuantity := 0

	for _, dc := range *domainDCs {
		if dc.ProductQuantity > maxQuantity {
			maxQuantity = dc.ProductQuantity
			maxDC = &dc
		}
	}

	if maxDC == nil {
		return nil, fmt.Errorf("could not determine distribution center with highest product quantity")
	}

	return maxDC, nil
}
