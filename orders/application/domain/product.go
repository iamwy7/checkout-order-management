package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("product name is required")
	}
	if p.Price <= 0.00 {
		return errors.New("product invalid price")
	}
	return nil
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}
