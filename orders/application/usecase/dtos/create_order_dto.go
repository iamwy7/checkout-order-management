package dtos

type CreateOrderInputDTO struct {
	Products []ProductInputDTO `json:"products"`
	Zone     string            `json:"zone"` // chosing criteria
	State    string            `json:"state"`
}

type ProductInputDTO struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type CreatedOrderOutputDTO struct {
	Id            string             `json:"id"`
	Zone          string             `json:"zone"`
	State         string             `json:"state"`
	Status        string             `json:"status"`
	CreatedAt     string             `json:"created_at"`
	UpdatedAt     string             `json:"updated_at"`
	Products      []ProductOutputDTO `json:"products"`
	ProductsCount int                `json:"products_count"`
}

type ProductOutputDTO struct {
	Id                 string                      `json:"id"`
	Name               string                      `json:"name"`
	Price              float64                     `json:"price"`
	Quantity           int                         `json:"quantity"`
	DistributionCenter DistributionCenterOutputDTO `json:"distribution_center"`
}

type DistributionCenterOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Zone string `json:"zone"`
}
