package domain

import "github.com/iamwy7/meli-challenge/orders/application/shared"

type DistributionCenterInterface interface {
	Validate() error
}

type DistributionCenter struct {
	Id              string
	Name            string
	Zone            shared.Zone
	State           string
	Status          shared.Status
	ProductQuantity int
}

func NewDistributionCenter(id string, name string, zone shared.Zone, state string, status shared.Status, quantity int) (*DistributionCenter, error) {
	dc := &DistributionCenter{
		Id:              id,
		Name:            name,
		Zone:            zone,
		State:           state,
		Status:          status,
		ProductQuantity: quantity,
	}
	if err := dc.Validate(); err != nil {
		return nil, err
	}
	return dc, nil
}

func (dc *DistributionCenter) Validate() error {
	if dc.Id == "" {
		return ErrDCInvalidId
	}
	if dc.Name == "" {
		return ErrDCNameRequired
	}
	if dc.Zone == "" {
		return ErrDCZoneRequired
	}
	if len(dc.Zone) > 2 {
		return ErrDCInvalidZone
	}
	if dc.ProductQuantity <= 0 {
		return ErrDCInvalidQuantity
	}
	return nil
}
