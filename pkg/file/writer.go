package file

import (
	"io/ioutil"

	"github.com/go-openapi/spec"

	"github.com/far4599/swagger-openapiv2-merge/pkg/marshaller"
)

type writer struct {
	filePath string
	format   marshaller.OutputFormat
}

func NewFileWriter(filePath string, format marshaller.OutputFormat) *writer {
	return &writer{
		filePath: filePath,
		format:   format,
	}
}

func (w writer) Write(v *spec.Swagger) error {
	byteValue, err := marshaller.Marshal(v, w.format)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(w.filePath, byteValue, 0600)
}
