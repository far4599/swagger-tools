package file

import (
	"errors"
)

var (
	ErrInvalidJSONFormat = errors.New("invalid json format")
	ErrInvalidYAMLFormat = errors.New("invalid yaml format")
	ErrInvalidArgument   = errors.New("invalid argument")
	ErrDirNotExist       = errors.New("directory does not exist")
	ErrEmptyDir          = errors.New("empty directory")
)
