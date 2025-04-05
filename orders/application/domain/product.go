package domain

type ProductInterface interface {
	Validate() error
	NewProduct(name string, price float64) (*Product, error)
}

type ProductStatus string

const (
	ACTIVE   ProductStatus = "ACTIVE"
	INACTIVE ProductStatus = "APPROVED"
	DELETED  ProductStatus = "DELETED"
)

type Product struct {
	Id                 string //UUID
	Name               string
	Price              float64
	Quantity           int
	Status             ProductStatus // idea: worker that deactivates products
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
	if p.Status != ACTIVE {
		return ErrProductIsnActive
	}
	return nil
}

func NewProduct(id string, name string, price float64, quantity int, status ProductStatus, dc DistributionCenter) (*Product, error) {
	product := &Product{
		Id:                 id,
		Name:               name,
		Price:              price,
		Quantity:           quantity,
		Status:             status,
		DistributionCenter: dc,
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}
