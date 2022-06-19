package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
)

type dirReader struct {
	dirPath string

	optWithSubdir bool
	optFilterExt  string
}

func NewDirReader(dirPath string) *dirReader {
	return &dirReader{
		dirPath: dirPath,
	}
}

func (r dirReader) Read(out *[]*spec.Swagger) error {
	if out == nil {
		return errors.Wrapf(ErrInvalidArgument, "out must be a slice, given: nil")
	}
	if *out == nil {
		*out = make([]*spec.Swagger, 0)
	}

	files, err := ioutil.ReadDir(r.dirPath)
	if errors.Is(err, os.ErrNotExist) {
		return ErrDirNotExist
	}
	if err != nil {
		return errors.Wrapf(err, "failed to get directory '%s' content", r.dirPath)
	}

	if len(files) == 0 {
		return ErrEmptyDir
	}

	for _, f := range files {
		var loopErr error

		filename := path.Join(r.dirPath, f.Name())

		if f.IsDir() {
			if r.optWithSubdir {
				loopErr = NewDirReader(filename).
					WithSubdir(r.optWithSubdir).
					WithExtFilter(r.optFilterExt).
					Read(out)
				if loopErr != nil {
					return loopErr
				}
			}

			continue
		}
		if !strings.HasSuffix(filename, r.optFilterExt) {
			continue
		}

		var newSpec spec.Swagger
		loopErr = NewFileReader(filename).Read(&newSpec)
		if loopErr != nil {
			fmt.Printf("failed to read swagger spec from file '%s': '%v'", filename, loopErr)
		}

		*out = append(*out, &newSpec)
	}

	return nil
}

func (r *dirReader) WithSubdir(v bool) *dirReader {
	r.optWithSubdir = v
	return r
}

func (r *dirReader) WithExtFilter(filterExt string) *dirReader {
	r.optFilterExt = filterExt
	return r
}
