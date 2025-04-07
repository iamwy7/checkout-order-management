package usecase

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
	"github.com/iamwy7/meli-challenge/orders/application/usecase/dtos"
)

func MapOrderToOrderOutputDto(order *domain.Order) *dtos.CreatedOrderOutputDto {
	productsDto := make([]dtos.ProductOutputDto, len(order.Products))
	for _, product := range order.Products {
		dc := dtos.DistributionCenterOutputDto{
			Id:   product.DistributionCenter.Id,
			Name: product.DistributionCenter.Name,
			Zone: string(product.DistributionCenter.Zone),
		}
		productDto := dtos.ProductOutputDto{
			Id:                 product.Id,
			Name:               product.Name,
			Price:              product.Price,
			Quantity:           product.Quantity,
			DistributionCenter: dc,
		}
		productsDto = append(productsDto, productDto)
	}

	orderDto := &dtos.CreatedOrderOutputDto{
		Id:        order.Id,
		Zone:      string(order.Zone),
		State:     order.State,
		Status:    string(order.Status),
		CreatedAt: order.CreatedAt.String(),
		UpdatedAt: order.UpdatedAt.String(),
		Products:  productsDto,
	}
	return orderDto
}
