package marshaller

import (
	"errors"
)

var (
	ErrInvalidJSONFormat = errors.New("invalid json format")
	ErrInvalidYAMLFormat = errors.New("invalid yaml format")
)
