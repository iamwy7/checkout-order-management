package api_handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	custom_errors "github.com/iamwy7/meli-challenge/distribution_centers/errors"
	"github.com/iamwy7/meli-challenge/distribution_centers/repo"
)

// DistributionCentersHandler is our principal api handler
type DistributionCentersHandler struct {
	dcRepository *repo.DistributionCenterDB
}

// NewDistributionCentersHandler creates an instance of DistributionCentersHandler
func NewDistributionCentersHandler(dcRepository *repo.DistributionCenterDB) *DistributionCentersHandler {
	return &DistributionCentersHandler{dcRepository: dcRepository}
}

// GetDistributionCenter handles the request to get the distribution centers of an itemId
//
//	@Summary		List distribution centers
//	@Description	Get all distribution centers of an itemId
//	@Tags			DistributionCenters
//	@Accept			json
//	@Produce		json
//	@Param			itemId	query		string	true	"itemId"
//	@Success		200		{object}	dtos.DistributionCenters
//	@Failue			404 	{object} 	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/distributioncenters [get]
func (dch *DistributionCentersHandler) GetDistributionCenter(w http.ResponseWriter, r *http.Request) {
	itemId := r.URL.Query().Get("itemId")
	responseDcs, err := dch.dcRepository.GetDistributionCentersByProductId(itemId)
	if errors.Is(err, custom_errors.ErrProductInvalidId) {
		dch.writeErrorResponse(w, fmt.Sprintf("%v for itemId:%v", err.Error(), itemId), http.StatusNotFound)
		return
	}
	if err != nil {
		dch.writeErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseDcs)
}

// writeErrorResponse writes an error response in JSON format
func (dch *DistributionCentersHandler) writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	log.Printf("handler error: %v", message)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Message string `json:"message" example:" message about the error"`
}
