package error

import "errors"

var (
	ErrNoConnection          = errors.New("No connection set")
	ErrMissingRequiredFilter = errors.New("Required fields are not set")
	ErrModelNotLoaded        = errors.New("Model was not initialized")
	ErrCreateTransaction     = errors.New("Transaction cant be created")
)
