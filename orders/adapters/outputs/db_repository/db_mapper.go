package db_repository

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
	"github.com/iamwy7/meli-challenge/orders/application/shared"
)

func MapOrderDtoToOrder(orderDto OrderDto) (*domain.Order, error) {
	orderProducts := make([]domain.Product, len(orderDto.Products))
	for _, productDto := range orderDto.Products {
		product, err := MapProductDtoToProduct(productDto)
		if err != nil {
			return nil, err
		}
		orderProducts = append(orderProducts, *product)
	}

	order := &domain.Order{
		Id:       orderDto.Id,
		Zone:     shared.Zone(orderDto.Zone),
		State:    orderDto.State,
		Status:   shared.Status(orderDto.Status),
		Products: orderProducts,
	}

	return order, nil
}

func MapProductDtoToProduct(productDto ProductDto) (*domain.Product, error) {
	productDC, err := MapDistributionCenterDtoToDistributionCenter(productDto.DistributionCenter)
	if err != nil {
		return nil, err
	}
	product, err := domain.NewProduct(productDto.Id, productDto.Name, productDto.Price, productDto.Quantity, *productDC)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func MapDistributionCenterDtoToDistributionCenter(dcDto DistributionCenterDto) (*domain.DistributionCenter, error) {
	domainDC, err := domain.NewDistributionCenter(dcDto.Id, dcDto.Name, shared.Zone(dcDto.Zone), dcDto.State, shared.Status(dcDto.Status), dcDto.ProductQuantity)
	if err != nil {
		return nil, err
	}
	return domainDC, nil
}
