package utils

import (
	"reflect"

	"github.com/getkin/kin-openapi/openapi3"
)

// ExtractSecurityRequirementKey retrieves the initial security key
func ExtractSecurityRequirementKey(input openapi3.SecurityRequirement) string {
	for key := range input {
		return key
	}

	return ""
}

// ExtractInterfaceOfMap retrieves interface-based data with structure of hashmap
func ExtractInterfaceOfMap(data interface{}) (bool, map[string]interface{}) {
	result := make(map[string]interface{})
	m := reflect.ValueOf(data)

	if m.Kind() != reflect.Map {
		return false, nil
	} else {
		for _, key := range m.MapKeys() {
			value := m.MapIndex(key)

			switch value.Interface().(type) {
			case string:
				result[key.Interface().(string)] = value.Interface().(string)
			case bool:
				result[key.Interface().(string)] = value.Interface().(bool)
			case map[string]interface{}:
				result[key.Interface().(string)] = value.Interface().(map[string]interface{})
			default:
				result[key.Interface().(string)] = value.Interface()
			}

		}
	}

	return true, result
}

// IsDataExistInArray works to check whether particular key exists on array or not
func IsDataExistInArray(array interface{}, key interface{}) bool {
	arr := reflect.ValueOf(array)

	if arr.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == key {
			return true
		}
	}

	return false
}
