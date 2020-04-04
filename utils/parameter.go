package utils

import (
	"github.com/getkin/kin-openapi/openapi3"
)

const (
	// ParamHeader is header
	ParamHeader = "header"
	// ParamQuery is query
	ParamQuery = "query"
)

// IsParameterExist works to check the existency of particular param type on list of parameter
func IsParameterExist(params openapi3.Parameters, paramType string) bool {
	if len(params) > 0 {
		for _, param := range params {
			if param.Value.In == paramType {
				return true
			}
		}
	}

	return false
}
