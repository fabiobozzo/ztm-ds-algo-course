package constants

import "errors"

var (
	ErrNotFound         = errors.New("item not found")
	ErrEmpty            = errors.New("the data structure is empty")
	ErrIndexOutOfBounds = errors.New("index out of bounds")
	ErrInvalidOperation = errors.New("invalid operation")
)
