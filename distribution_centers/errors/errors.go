package custom_errors

import "errors"

var (
	ErrProductInvalidId = errors.New("product id is invalid")
	ErrUnexpectedError  = errors.New("unexpected error occurred")
)
