package dtos

// CreateOrderInputDTO represents the input data required to create an order.
type CreateOrderInputDTO struct {
	Products []ProductInputDTO `json:"products"`
	Zone     string            `json:"zone" example:"S1"` // choosing criteria
	State    string            `json:"state" example:"SP"`
}

// ProductInputDTO represents the input data required to get an order.
type ProductInputDTO struct {
	Id       string  `json:"id" example:"2b494226-db60-4dd0-b2a0-141b3ae52517"`
	Name     string  `json:"name" example:"Product A"`
	Price    float64 `json:"price" example:"10.5"`
	Quantity int     `json:"quantity" example:"4"`
}

// CreatedOrderOutputDTO represents the output data after creating an order.
type CreatedOrderOutputDTO struct {
	Id            string             `json:"id" example:"adec01f0-0393-436d-a441-04575e6624bb"`
	Zone          string             `json:"zone" example:"S1"`
	State         string             `json:"state" example:"SP"`
	Status        string             `json:"status" example:"PENDING"`
	CreatedAt     string             `json:"created_at" example:"2025-04-07T22:00:00Z"`
	UpdatedAt     string             `json:"updated_at" example:"2025-04-07T22:30:00Z"`
	Products      []ProductOutputDTO `json:"products"`
	ProductsCount int                `json:"products_count" example:"3"`
}

// ProductOutputDTO represents the output data of the products of an order.
type ProductOutputDTO struct {
	Id                 string                      `json:"id" example:"3195b102-353a-4d84-9639-7eda62ef5f08"`
	Name               string                      `json:"name" example:"Product A"`
	Price              float64                     `json:"price" example:"10.5"`
	Quantity           int                         `json:"quantity" example:"2"`
	DistributionCenter DistributionCenterOutputDTO `json:"distribution_center"`
}

// DistributionCenterOutputDTO represents the output data of the distribution center of a product.
type DistributionCenterOutputDTO struct {
	Id   string `json:"id" example:"a2d705c1-7431-456b-b65f-bee431694b46"`
	Name string `json:"name" example:"DC1"`
	Zone string `json:"zone" example:"S1"`
}
