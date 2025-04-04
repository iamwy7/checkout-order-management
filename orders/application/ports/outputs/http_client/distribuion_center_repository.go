package ports

import (
	"github.com/iamwy7/meli-challenge/orders/application/domain"
)

type DistributionCentersRepository interface {
	GetDCsByItemId(itemID string) (*[]domain.DistributionCenter, error)
}
