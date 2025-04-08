package http_client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/iamwy7/meli-challenge/orders/application/domain"
	out_ports "github.com/iamwy7/meli-challenge/orders/application/ports/outputs"
)

type DistributionCenterAdapter struct {
	BaseURL string
}

func NewDistributionCenterAdapterFactory(baseUrl string) out_ports.DistributionCentersRepository {
	return &DistributionCenterAdapter{BaseURL: baseUrl}
}

func (dca *DistributionCenterAdapter) GetDCsByItemId(itemId string) (*[]domain.DistributionCenter, error) {
	path := fmt.Sprintf("%s/distributioncenters?itemId=%s&zone=%s&status=%s", dca.BaseURL, itemId, "SP", "ACTIVE")

	// Prep request with path
	httpReq, _ := http.NewRequest("GET", path, nil)

	// Get the current usable http_client and do the request
	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		log.Printf("integration error to get distribution centers: %v", err.Error())
		return nil, ErrIntegrationServer
	}

	defer func() {
		if httpResp.Body != nil {
			httpResp.Body.Close()
		}
	}()

	if httpResp.StatusCode != http.StatusOK {
		log.Printf("failed to get distribution centers by id: %v with status code: %d", itemId, httpResp.StatusCode)
		return nil, ErrIntegrationClient
	}

	// Decode the response
	var dcResp DistributionCenterResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&dcResp); err != nil {
		log.Printf("failed to decode response: %v", err.Error())
		return nil, ErrIntegrationDecodeJson
	}

	// Check if the response is empty
	if len(dcResp.DistributionCenters) == 0 {
		log.Printf("no distribution centers found for itemId: %v", itemId)
		return nil, ErrIntegrationDCsEmpty
	}

	domainDCs, err := MapDCResponseToDomain(&dcResp)
	if err != nil {
		log.Printf("failed to map response to domain: %v", err.Error())
		return nil, ErrIntegrationServer
	}

	return domainDCs, nil
}
