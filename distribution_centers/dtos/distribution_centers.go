package dtos

type DistributionCenter struct {
	Id       string `json:"id" example:"a2d705c1-7431-456b-b65f-bee431694b46"`
	Name     string `json:"name" example:"CD1"`
	Zone     string `json:"zone" example:"W1"`
	State    string `json:"state" example:"SP"`
	Status   string `json:"status" example:"ACTIVE"`
	Quantity int    `json:"quantity" example:"633"`
}

type DistributionCenters struct {
	ItemId             string               `json:"itemId" example:"493d60ac-64c4-4bc6-8fbf-eefe752790c4"`
	DistributionCenter []DistributionCenter `json:"distributionCenters"`
}

func NewDistributionCenters(itemId string, dcs []DistributionCenter) *DistributionCenters {
	return &DistributionCenters{
		ItemId:             itemId,
		DistributionCenter: dcs,
	}
}
