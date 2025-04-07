package db_repository

type OrderDto struct {
	Id       string // UUID
	Zone     string
	State    string
	Status   string
	Products []ProductDto
}

type ProductDto struct {
	Id                 string //UUID
	Name               string
	Price              float64
	Quantity           int
	Status             string
	DistributionCenter DistributionCenterDto
}

type DistributionCenterDto struct {
	Id              string
	Name            string
	Zone            string
	State           string
	Status          string
	ProductQuantity int
}
