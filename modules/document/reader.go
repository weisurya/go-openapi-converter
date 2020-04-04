package moduledoc

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/unidoc/unioffice/document"
)

// ReadDocTemplate to use existing document template
func ReadDocTemplate(path string) (*document.Document, error) {
	doc, err := document.Open(path)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// ReadOpenAPI to read OpenAPI spec
func ReadOpenAPI(path string) (*openapi3.Swagger, error) {
	spec, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(path)
	if err != nil {
		return nil, err
	}

	return spec, nil
}
