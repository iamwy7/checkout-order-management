package http_client

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
)

func MapDCResponseToDomain(dcResp *DistributionCenterResponse) (*[]domain.DistributionCenter, error) {
	domainDCs := make([]domain.DistributionCenter, len(dcResp.DistributionCenters))

	for i, dc := range dcResp.DistributionCenters {
		tempDC, err := domain.NewDistributionCenter(dc.Id, dc.Name, dc.Zone, dc.State, dc.Status, dc.Quantity)
		if err != nil {
			return nil, err
		}
		domainDCs[i] = *tempDC

	}
	return &domainDCs, nil
}
