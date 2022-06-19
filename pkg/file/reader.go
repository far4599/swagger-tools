package file

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type fileReader struct {
	filePath string
}

func NewFileReader(filePath string) *fileReader {
	return &fileReader{
		filePath: filePath,
	}
}

func (r fileReader) Read(out *spec.Swagger) error {
	byteValue, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	return unmarshal(byteValue, out)
}

func isJSON(content []byte) bool {
	return json.Unmarshal(content, new(json.RawMessage)) == nil
}

func unmarshal(content []byte, out interface{}) error {
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
