package http_client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type DistributionCenterAdapter struct {
	BaseURL string
}

func (dca *DistributionCenterAdapter) GetDCsByItemId(itemID string) (*DistributionCenterResponse, error) {
	path := fmt.Sprintf("%s/distributioncenters?itemId=%s", dca.BaseURL, itemID)

	// Prep request with path
	httpReq, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Fatalf("failed to prep request to get distribution centers for reason:%v", err.Error())
		return nil, err
	}

	// Get the current usable http_client and do the request
	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		log.Fatalf("integration error to get distribution centers with status: %d and error %v", httpResp.StatusCode, err.Error())
		return nil, ErrIntegrationServer
	}

	// Close when done
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		log.Fatalf("failed to get distribution centers by id: %v with status code: %d", itemID, httpResp.StatusCode)
		return nil, ErrIntegrationClient
	}

	// Decode the response
	var dcResp DistributionCenterResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&dcResp); err != nil {
		log.Fatalf("failed to decode response: %v", err.Error())
		return nil, ErrIntegrationDecodeJson
	}
	// Check if the response is empty (that definetly means that the itemID is not valid or something else)
	if len(dcResp.DistributionCenters) == 0 {
		log.Fatalf("no distribution centers found for itemID: %v", itemID)
		return nil, ErrIntegrationDCsEmpty
	}
	return &dcResp, nil
}
