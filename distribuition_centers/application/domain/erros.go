package domain

import "errors"

var (
	ErrDCEmpty = errors.New("distribution center has no products")
	ErrSPEmpty = errors.New("distribution center doesnt have this product to deliver")
)
