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

type Order struct {
	Id       string // UUID
	Products []Product
	Zone     Zone
	State    string
}

func (o *Order) Validate() error {
	if len(o.Products) == 0 {
		return errors.New("order must have some products to proccess")
	}
	if len(o.Products) > 100 {
		return errors.New("order have too many items, 100 is the limit")
	}
	if o.Zone == "" {
		return errors.New("order zone is required")
	}
	if len(o.Zone) != 2 {
		return errors.New("order zone must be 2 characters long")
	}
	return nil
}
