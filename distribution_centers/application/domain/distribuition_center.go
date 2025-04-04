package domain

import "errors"

type Zone string

const (
	North1 Zone = "N1"
	North2 Zone = "N2"
	East1  Zone = "E1"
	East2  Zone = "E2"
	West   Zone = "W1"
	Center Zone = "C1"
	South1 Zone = "S1"
	South2 Zone = "S2"
)

type DistributionCenter struct {
	Id      string // UUID
	Name    string // CD1, CD2 etc...
	Zone    Zone
	Storage []StoredProduct
}

func (cd *DistributionCenter) Validate() error {
	if cd.Name == "" {
		return errors.New("distribution center name is required")
	}
	if len(cd.Name) > 4 {
		return errors.New("distribution center must be a maximum of 4 characters")
	}
	if cd.Zone == "" {
		return errors.New("distribution center zone is required")
	}
	if len(cd.Name) != 2 {
		return errors.New("distribution center zone must be 2 characters long")
	}
	return nil
}
func NewDistributionCenter(name string, zone Zone) (*DistributionCenter, error) {
	return nil, nil
}

func (dc *DistributionCenter) AddProductToStorage(newProduct StoredProduct) (*StoredProduct, error) {
	newStoredProduct, err := NewStoredProduct(dc, newProduct.Name, newProduct.Price, newProduct.Quantity)
	if err != nil {
		return nil, err
	}
	dc.Storage = append(dc.Storage, *newStoredProduct)
	return newStoredProduct, nil
}
