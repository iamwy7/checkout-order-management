package http_client

type EachDistributionCenter struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Zone     string `json:"zone"`
	State    string `json:"state"`
	Status   string `json:"status"`
	Quantity int    `json:"quantity"`
}

type DistributionCenterResponse struct {
	ItemId              string                   `json:"itemId"`
	DistributionCenters []EachDistributionCenter `json:"distributionCenters"`
}
