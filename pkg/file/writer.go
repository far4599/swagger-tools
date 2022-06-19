package file

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-openapi/spec"
	"gopkg.in/yaml.v3"
)

type outputFormat uint

const (
	JSONFormat outputFormat = iota
	YAMLFormat
)

type writer struct {
	filePath string
	format   outputFormat
}

func NewFileWriter(filePath string, format outputFormat) *writer {
	return &writer{
		filePath: filePath,
		format:   format,
	}
}

func (w writer) Write(v *spec.Swagger) error {
	byteValue, err := marshal(v, w.format)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(w.filePath, byteValue, 0644)
}

func marshal(v interface{}, format outputFormat) ([]byte, error) {
	if format == YAMLFormat {
		return yaml.Marshal(v)
	}

	return json.MarshalIndent(v, "", "  ")
}
