package http_client

import "github.com/iamwy7/meli-challenge/orders/application/domain"

func MapResponseToDomain(dcResp *DistributionCenterResponse) *[]domain.DistributionCenter {
	domainDCs := make([]domain.DistributionCenter, len(dcResp.DistributionCenters))
	for i, dc := range dcResp.DistributionCenters {
		domainDCs[i] = domain.DistributionCenter{
			ID:       dc.ID,
			Name:     dc.Name,
			Zone:     dc.Zone,
			Quantity: dc.Quantity,
		}
	}
	return &domainDCs
}
