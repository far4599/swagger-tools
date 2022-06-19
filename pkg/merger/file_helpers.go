package merger

import (
	"github.com/go-openapi/spec"

	"github.com/far4599/swagger-openapiv2-merge/pkg/file"
)

func GetSpecFromFile(filename string) (*spec.Swagger, error) {
	var item spec.Swagger
	err := file.NewFileReader(filename).Read(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func WriteSpecToFile(base *spec.Swagger, filename string) error {
	return file.NewFileWriter(filename, file.JSONFormat).Write(base)
}

func LoadSpecsFromDir(dir, filterExt string, withSubdir bool) ([]*spec.Swagger, error) {
	result := make([]*spec.Swagger, 0)
	err := file.NewDirReader(dir).
		WithSubdir(withSubdir).
		WithExtFilter(filterExt).
		Read(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
