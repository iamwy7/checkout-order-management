package domain

import "errors"

var (
	// order errors
	ErrOrderInvalidId           = errors.New("order doenst exist")
	ErrOrderWithoutProducts     = errors.New("order must have some products to proccess")
	ErrOrderWithTooMuchProducts = errors.New("order have too many items, 100 is the limit")
	ErrOrderZoneRequired        = errors.New("order zone is invalid at the moment")
	ErrOrderInvalidState        = errors.New("order state is invalid")
	ErrOrderInvalidZone         = errors.New("order zone is invalid")

	// products errors
	ErrProductInvalidId        = errors.New("product id is invalid")
	ErrProductInvalidName      = errors.New("product name is required")
	ErrProductInvalidPrice     = errors.New("product invalid price")
	ErrProductWithoutDC        = errors.New("product distribution center is required")
	ErrProductInvalidQuantity  = errors.New("product quantity is required")
	ErrProductIsnActive        = errors.New("product is no longer active")
	ErrProductQuantityExceeded = errors.New("product quantity exceeded the limit of the distribution center")

	// distribution center errors
	ErrDCInvalidId       = errors.New("distribution center id is required")
	ErrDCNameRequired    = errors.New("distribution center name is required")
	ErrDCZoneRequired    = errors.New("distribution center zone is required")
	ErrDCInvalidZone     = errors.New("distribution center zone is invalid")
	ErrDCInvalidQuantity = errors.New("distribution center product quantity must be greater than 0")

	ErrUnexpectedError = errors.New("unexpected error occurred")
)
