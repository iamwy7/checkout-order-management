package domain

type DistributionCenterInterface interface {
	Validate() error
}

type DistributionCenter struct {
	Id              string
	Name            string
	Zone            Zone
	ProductQuantity int
}

func NewDistributionCenter(id string, name string, zone Zone, quantity int) (*DistributionCenter, error) {
	dc := &DistributionCenter{
		Id:              id,
		Name:            name,
		Zone:            zone,
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
