package marshaller

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type OutputFormat uint

const (
	JSONFormat OutputFormat = iota
	YAMLFormat
)

func Marshal(v interface{}, format OutputFormat) ([]byte, error) {
	if format == YAMLFormat {
		return yaml.Marshal(v)
	}

	return json.MarshalIndent(v, "", "  ")
}
