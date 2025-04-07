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
	Id        string // UUID
	Zone      shared.Zone
	State     string
	Status    shared.Status
	CreatedAt time.Time
	UpdatedAt time.Time
	Products  []Product
}

func (o *Order) Validate() error {
	if len(o.Products) == 0 {
		return ErrOrderWithoutProducts
	}
	if len(o.Products) > 100 {
		return ErrOrderWithTooMuchProducts
	}
	if o.State != "SP" {
		return ErrOrderZoneRequired
	}
	if o.Zone == "" {
		return ErrOrderZoneRequired
	}
	if len(o.Zone) != 2 {
		return ErrOrderInvalidZone
	}
	return nil
}
func NewOrder(products []Product, zone string, state string) (*Order, error) {
	validated_zone, err := shared.ValidateZone(zone)
	if err != nil {
		return nil, err
	}
	order := &Order{
		Id:        uuid.NewString(),
		Zone:      validated_zone,
		State:     state,
		Status:    shared.PENDING,
		CreatedAt: time.Now(), // example: Date.Format("2025-01-02 15:04:05")
		UpdatedAt: time.Now(), // example: Date.Format("2025-01-02 15:04:05")
		Products:  products,
	}
	if err := order.Validate(); err != nil {
		return nil, err
	}
	return order, nil
}
