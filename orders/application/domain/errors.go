package domain

import "errors"

var (
	// order errors
	ErrOrderWithoutProducts     = errors.New("order must have some products to proccess")
	ErrOrderWithTooMuchProducts = errors.New("order have too many items, 100 is the limit")
	ErrOrderZoneRequired        = errors.New("order zone is required")
	ErrOrderInvalidZone         = errors.New("order zone is required")

	// products errors
	ErrProductInvalidName     = errors.New("product name is required")
	ErrProductInvalidPrice    = errors.New("product invalid price")
	ErrProductWithoutDC       = errors.New("product distribution center is required")
	ErrProductInvalidQuantity = errors.New("product quantity is required")
	ErrProductIsnActive       = errors.New("product is no longer active")

	// distribution center errors
	ErrDCInvalidId       = errors.New("distribution center id is required")
	ErrDCNameRequired    = errors.New("distribution center name is required")
	ErrDCZoneRequired    = errors.New("distribution center zone is required")
	ErrDCInvalidZone     = errors.New("distribution center zone must be 2 characters long")
	ErrDCInvalidQuantity = errors.New("distribution center product quantity must be greater than 0")
)
