package file

import (
	"io/ioutil"

	"github.com/go-openapi/spec"

	"github.com/far4599/swagger-openapiv2-merge/pkg/marshaller"
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

	return marshaller.Unmarshal(byteValue, out)
}
