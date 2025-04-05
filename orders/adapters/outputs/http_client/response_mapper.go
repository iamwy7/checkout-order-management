package http_client

import "github.com/iamwy7/meli-challenge/orders/application/domain"

func MapResponseToDomain(dcResp *DistributionCenterResponse) *[]domain.DistributionCenter {
	// TODO
	return nil
}

func validateZone(zone string) (*domain.Zone, error) {
	if len(zone) != 2 {
		return nil, domain.ErrDCInvalidZone
	}
	switch zone {
	case string(domain.North1), string(domain.North2), string(domain.East1), string(domain.East2), string(domain.West), string(domain.Center), string(domain.South1), string(domain.South2):
		zoneValue := domain.Zone(zone)
		return &zoneValue, nil
	default:
		return nil, domain.ErrOrderInvalidZone
	}
}
