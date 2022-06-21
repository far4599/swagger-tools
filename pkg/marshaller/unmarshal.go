package marshaller

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func isJSON(content []byte) bool {
	return json.Unmarshal(content, new(json.RawMessage)) == nil
}

func Unmarshal(content []byte, out interface{}) error {
	if isJSON(content) {
		err := json.Unmarshal(content, out)
		if err != nil {
			return errors.Wrapf(ErrInvalidJSONFormat, "error: %v", err)
		}
	} else {
		err := yaml.Unmarshal(content, out)
		if err != nil {
			return errors.Wrapf(ErrInvalidYAMLFormat, "error: %v", err)
		}
	}

	return nil
}
