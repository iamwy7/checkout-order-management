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

func NewDistributionCenter(id string, name string, zone string, state string, status string, quantity int) (*DistributionCenter, error) {
	validatedStatus := shared.CheckStatus(status)
	if validatedStatus == "" {
		return nil, ErrDcInvalidStatus
	}
	validatedZone := shared.CheckZone(zone)
	if validatedZone == "" {
		return nil, ErrDCInvalidZone
	}
	dc := &DistributionCenter{
		Id:              id,
		Name:            name,
		Zone:            validatedZone,
		State:           state,
		Status:          validatedStatus,
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
	if dc.ProductQuantity <= 0 {
		return ErrDCInvalidQuantity
	}
	return nil
}
