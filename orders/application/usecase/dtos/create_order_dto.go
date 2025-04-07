package dtos

type CreateOrderInputDTO struct {
	Products []ProductInputDto `json:"products"`
	Zone     string            `json:"zone"` // chosing criteria
	State    string            `json:"state"`
}

type ProductInputDto struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type CreatedOrderOutputDto struct {
	Id            string             `json:"id"`
	Zone          string             `json:"zone"`
	State         string             `json:"state"`
	Status        string             `json:"status"`
	CreatedAt     string             `json:"created_at"`
	UpdatedAt     string             `json:"updated_at"`
	Products      []ProductOutputDto `json:"products"`
	ProductsCount int                `json:"products_count"`
}

type ProductOutputDto struct {
	Id                 string                      `json:"id"`
	Name               string                      `json:"name"`
	Price              float64                     `json:"price"`
	Quantity           int                         `json:"quantity"`
	DistributionCenter DistributionCenterOutputDto `json:"distribution_center"`
}

type DistributionCenterOutputDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Zone string `json:"zone"`
}
