package merger

import (
	"fmt"
	"reflect"

	"github.com/go-openapi/spec"
)

type merger struct {
	result                *spec.Swagger
	isInfoOverrideAllowed bool
}

func NewMerger(specVersion string) *merger {
	m := &merger{
		isInfoOverrideAllowed: true,
	}

	m.result = new(spec.Swagger)
	m.result.Swagger = specVersion

	return m
}

func (m *merger) MergeInfo(info *spec.Swagger, isInfoOverrideAllowed bool) error {
	if info == nil {
		return ErrNoSpecProvided
	}

	m.result.Info = info.Info
	m.result.Host = info.Host
	m.isInfoOverrideAllowed = isInfoOverrideAllowed

	return nil
}

func (m merger) Result() *spec.Swagger {
	return m.result
}

func (m *merger) MergeSpec(newSpec *spec.Swagger) {
	if m.isInfoOverrideAllowed && newSpec.Info != nil {
		m.result.Info = newSpec.Info
	}

	if newSpec.ExternalDocs != nil {
		m.result.ExternalDocs = newSpec.ExternalDocs
	}

	m.result.Consumes = uniqueStrings(append(m.result.Consumes, newSpec.Consumes...))
	m.result.Produces = uniqueStrings(append(m.result.Produces, newSpec.Produces...))
	m.result.Schemes = uniqueStrings(append(m.result.Schemes, newSpec.Schemes...))
	m.result.Tags = uniqueTags(append(m.result.Tags, newSpec.Tags...))

	if newSpec.Paths != nil {
		for key, value := range newSpec.Paths.Paths {
			if m.result.Paths == nil || len(m.result.Paths.Paths) == 0 {
				m.result.Paths = newSpec.Paths
				break
			}

			if path, ok := m.result.Paths.Paths[key]; ok && !reflect.DeepEqual(path, value) {
				fmt.Printf("path '%s' already exist and will be overwritten\n", key)
			}

			m.result.Paths.Paths[key] = value
		}
	}

	for key, value := range newSpec.Definitions {
		if m.result.Definitions == nil {
			m.result.Definitions = make(spec.Definitions)
		}

		if def, ok := m.result.Definitions[key]; ok && !reflect.DeepEqual(def, value) {
			fmt.Printf("definition '%s' already exist and will be overwritten\n", key)
		}

		m.result.Definitions[key] = value
	}

	for key, value := range newSpec.Parameters {
		if m.result.Parameters == nil {
			m.result.Parameters = make(map[string]spec.Parameter)
		}

		if param, ok := m.result.Parameters[key]; ok && !reflect.DeepEqual(param, value) {
			fmt.Printf("parameter '%s' already exist and will be overwritten\n", key)
		}

		m.result.Parameters[key] = value
	}

	for key, value := range newSpec.Responses {
		if m.result.Responses == nil {
			m.result.Responses = make(map[string]spec.Response)
		}

		if resp, ok := m.result.Responses[key]; ok && !reflect.DeepEqual(resp, value) {
			fmt.Printf("response '%s' already exist and will be overwritten\n", key)
		}

		m.result.Responses[key] = value
	}

	for key, value := range newSpec.SecurityDefinitions {
		if m.result.SecurityDefinitions == nil {
			m.result.SecurityDefinitions = make(spec.SecurityDefinitions)
		}

		if secDef, ok := m.result.SecurityDefinitions[key]; ok && !reflect.DeepEqual(secDef, value) {
			fmt.Printf("securityDefinition '%s' already exist and will be overwritten\n", key)
		}

		m.result.SecurityDefinitions[key] = value
	}

	for _, sec := range newSpec.Security {
		if len(newSpec.Security) == 0 {
			continue
		}

		if len(m.result.Security) == 0 {
			m.result.Security = newSpec.Security
			continue
		}

		for key, value := range sec {
			if sec, ok := m.result.Security[0][key]; ok && !reflect.DeepEqual(sec, value) {
				fmt.Printf("security '%s' already exist and will be overwritten\n", key)
			}

			m.result.Security[0][key] = value
		}
	}
}

func uniqueStrings(arr []string) []string {
	if arr == nil {
		return nil
	}

	m := make(map[string]struct{})
	for i := range arr {
		m[arr[i]] = struct{}{}
	}

	result := make([]string, 0, len(m))
	for s := range m {
		result = append(result, s)
	}

	return result
}

func uniqueTags(tags []spec.Tag) []spec.Tag {
	if tags == nil {
		return nil
	}

	m := make(map[string]spec.Tag)
	for i := range tags {
		m[tags[i].Name] = tags[i]
	}

	result := make([]spec.Tag, 0, len(m))
	for _, t := range m {
		result = append(result, t)
	}

	return result
}
