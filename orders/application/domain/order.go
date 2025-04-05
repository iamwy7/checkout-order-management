package domain

type OrderInterface interface {
	Validate() error
}

type OrderStatus string

const (
	PENDING  OrderStatus = "PENDING"
	APPROVED OrderStatus = "APPROVED"
	CANCELED OrderStatus = "CANCELED"
)

type Order struct {
	Id       string // UUID
	Products []Product
	Zone     Zone
	State    string
	Status   OrderStatus
}

func NewOrder(id string, products []Product, zone Zone, state string, status OrderStatus) (*Order, error) {
	order := &Order{
		Id:       id,
		Products: products,
		Zone:     zone,
		State:    state,
		Status:   status,
	}
	if err := order.Validate(); err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Validate() error {
	if len(o.Products) == 0 {	
		return ErrOrderWithoutProducts
	}
	if len(o.Products) > 100 {
		return ErrOrderWithTooMuchProducts
	}
	if o.Zone == "" {
		return ErrOrderZoneRequired
	}
	if len(o.Zone) != 2 {
		return ErrOrderInvalidZone
	}
	return nil
}
