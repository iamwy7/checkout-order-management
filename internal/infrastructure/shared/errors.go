package shared

import "errors"

var (
	// Status
	ErrInvalidStatus = errors.New("status is invalid")

	// Zone
	ErrInvalidZone = errors.New("zone is invalid")
)
