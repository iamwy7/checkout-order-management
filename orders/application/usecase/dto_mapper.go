package usecase

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
	"github.com/iamwy7/meli-challenge/orders/application/usecase/dtos"
)

func MapOrderToOrderOutputDto(order *domain.Order) *dtos.CreatedOrderOutputDTO {
	productsDto := make([]dtos.ProductOutputDTO, len(order.Products))
	for i, product := range order.Products {
		dc := dtos.DistributionCenterOutputDTO{
			Id:   product.DistributionCenter.Id,
			Name: product.DistributionCenter.Name,
			Zone: string(product.DistributionCenter.Zone),
		}
		productDto := dtos.ProductOutputDTO{
			Id:                 product.CatalogProductId,
			Name:               product.Name,
			Price:              product.Price,
			Quantity:           product.Quantity,
			DistributionCenter: dc,
		}
		productsDto[i] = productDto
	}

	orderDto := &dtos.CreatedOrderOutputDTO{
		Id:            order.Id,
		Zone:          string(order.Zone),
		State:         order.State,
		Status:        string(order.Status),
		CreatedAt:     order.CreatedAt.String(),
		UpdatedAt:     order.UpdatedAt.String(),
		Products:      productsDto,
		ProductsCount: order.ProductsCount,
	}
	return orderDto
}
