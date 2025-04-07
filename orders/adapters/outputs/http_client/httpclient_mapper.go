package http_client

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
	"github.com/iamwy7/meli-challenge/orders/application/shared"
)

func MapDCResponseToDomain(dcResp *DistributionCenterResponse) (*[]domain.DistributionCenter, error) {
	domainDcs := make([]domain.DistributionCenter, len(dcResp.DistributionCenters))

	for i, dc := range dcResp.DistributionCenters {
		dcZone, err := shared.ValidateZone(dc.Zone)
		if err != nil {
			return nil, err
		}
		dcStatus, err := shared.ValidateStatus(dc.Status)
		if err != nil {
			return nil, err
		}

		tempDC, err := domain.NewDistributionCenter(
			dc.Id,
			dc.Name,
			dcZone,
			dc.State,
			dcStatus,
			dc.Quantity,
		)
		if err != nil {
			return nil, err
		}
		domainDcs[i] = *tempDC

	}
	return &domainDcs, nil
}
