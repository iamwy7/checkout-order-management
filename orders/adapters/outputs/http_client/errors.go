package http_client

import "errors"

var (
	ErrIntegrationServer     = errors.New("integration error to get distribution centers, contact support")
	ErrIntegrationClient     = errors.New("failed to get distribution centers")
	ErrIntegrationDecodeJson = errors.New("failed to decode response")
	ErrIntegrationDCsEmpty   = errors.New("one or more products doenst have any distribution centers")
)
