package http_client

type EachDistributionCenter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Zone     string `json:"zone"`
	Quantity int    `json:"quantity"`
}

type DistributionCenterResponse struct {
	ItemID              string                   `json:"itemId"`
	DistributionCenters []EachDistributionCenter `json:"distributionCenters"`
}
