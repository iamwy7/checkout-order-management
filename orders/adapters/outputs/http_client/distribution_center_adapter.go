package http_client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/iamwy7/meli-challenge/orders/application/domain"
	ports "github.com/iamwy7/meli-challenge/orders/application/ports/outputs/http_client"
)

type DistributionCenterAdapter struct {
	BaseURL string
}

func NewDistributionCenterAdapterFactory(baseUrl string) ports.DistributionCentersRepository {
	return &DistributionCenterAdapter{BaseURL: baseUrl}
}

func (dca *DistributionCenterAdapter) GetDCsByItemId(itemId string) (*[]domain.DistributionCenter, error) {
	path := fmt.Sprintf("%s/distributioncenters?itemId=%s", dca.BaseURL, itemId)

	// Prep request with path
	httpReq, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Printf("failed to prep request to get distribution centers for reason:%v", err.Error())
		return nil, err
	}

	// Get the current usable http_client and do the request
	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		log.Printf("integration error to get distribution centers with status: %d and error %v", httpResp.StatusCode, err.Error())
		return nil, ErrIntegrationServer
	}

	// Close when done
	defer httpResp.Body.Close()
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

	// Check if the response is empty (that definetly means that the ItemId is not valid or something else)
	if len(dcResp.DistributionCenters) == 0 {
		log.Printf("no distribution centers found for itemId: %v", itemId)
		return nil, ErrIntegrationDCsEmpty
	}
	domainDCs := MapResponseToDomain(&dcResp)
	return domainDCs, nil
}
