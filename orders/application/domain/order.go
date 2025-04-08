package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/iamwy7/meli-challenge/orders/application/shared"
)

type OrderInterface interface {
	Validate() error
}

type Order struct {
	Id            string // UUID
	Zone          shared.Zone
	State         string
	Status        shared.Status
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Products      []Product
	ProductsCount int
}

func (o *Order) Validate() error {
	if len(o.Products) == 0 {
		return ErrOrderWithoutProducts
	}
	if len(o.Products) > 100 {
		return ErrOrderWithTooMuchProducts
	}
	if o.State != "SP" {
		return ErrOrderInvalidState
	}
	return nil
}
func NewOrder(products []Product, zone string, state string) (*Order, error) {
	validatedZone := shared.CheckZone(zone)
	if validatedZone == "" {
		return nil, ErrOrderInvalidZone
	}
	order := &Order{
		Id:            uuid.NewString(),
		Zone:          validatedZone,
		State:         state,
		Status:        shared.PENDING,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Products:      products,
		ProductsCount: len(products),
	}
	if err := order.Validate(); err != nil {
		return nil, err
	}
	return order, nil
}
