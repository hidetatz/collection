package collection

import "errors"

var (
	// ErrInvalidIndex indicates the given index is invalid for the operation.
	ErrInvalidIndex = errors.New("invalid index")
)
