package file

import (
	"errors"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrDirNotExist     = errors.New("directory does not exist")
	ErrEmptyDir        = errors.New("empty directory")
)
