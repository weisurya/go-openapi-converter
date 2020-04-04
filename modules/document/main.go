package moduledoc

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/unidoc/unioffice/document"
)

type docHandler struct {
	doc     *document.Document
	swagger *openapi3.Swagger
}

// NewDocumentHandler works as a handler to use document module
func NewDocumentHandler(doc *document.Document, swagger *openapi3.Swagger) DocRepository {
	return &docHandler{
		doc:     doc,
		swagger: swagger,
	}
}

// DocRepository list out available function
type DocRepository interface {
	BuildFrontpage()
	BuildTOC()
	BuildPaths()
}
