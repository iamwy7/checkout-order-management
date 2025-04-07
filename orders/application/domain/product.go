package domain

type ProductInterface interface {
	Validate() error
}

type Product struct {
	Id                 string //UUID
	Name               string
	Price              float64
	Quantity           int
	DistributionCenter DistributionCenter
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrProductInvalidName
	}
	if p.Price <= 0.00 {
		return ErrProductInvalidPrice
	}
	if p.DistributionCenter.Id == "" {
		return ErrProductWithoutDC
	}
	if p.Quantity <= 0 {
		return ErrProductInvalidQuantity
	}
	return nil
}

func NewProduct(id string, name string, price float64, quantity int, dc DistributionCenter) (*Product, error) {
	product := &Product{
		Id:                 id,
		Name:               name,
		Price:              price,
		Quantity:           quantity,
		DistributionCenter: dc,
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}
