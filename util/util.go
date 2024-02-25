package util

import (
	"fmt"
	"log"
	"strings"
)

func GetType(x interface{}) (dataType string, customType map[string]interface{}) {
	switch x.(type) {
	case int:
		dataType = "int"
	case int8:
		dataType = "int8"
	case int16:
		dataType = "int16"
	case int32:
		dataType = "int32"
	case int64:
		dataType = "int64"
	case uint:
		dataType = "uint"
	case uint8:
		dataType = "uint8"
	case uint16:
		dataType = "uint16"
	case uint32:
		dataType = "uint32"
	case uint64:
		dataType = "uint64"
	case float32:
		dataType = "float32"
	case float64:
		dataType = "float64"
	case string:
		dataType = "string"
	case bool:
		dataType = "bool"
	case []interface{}:
		dataType, customType = getSliceType(x.([]interface{}))
	case nil:
		dataType = "any"
	case map[string]interface{}:
		dataType = "object"
		customType = x.(map[string]interface{})
	default:
		log.Default().Panicf(fmt.Sprintf("%v is unknown", x))
	}
	return
}

func getSliceType(x []interface{}) (dataType string, customType map[string]interface{}) {
	if len(x) > 0 {
		t, c := GetType(x[0])
		return "[]" + t, c
	}
	return "[]interface{}", nil
}

func Capitalize(str string) string {
	if len(str) > 0 {
		stringSlice := strings.Split(str, "")
		return strings.ToUpper(stringSlice[0]) + strings.Join(stringSlice[1:], "")
	}
	return str
}

func ToCamelCase(str string) string {
	stringSlice := strings.Split(str, "_")
	result := ""
	for _, item := range stringSlice {
		result += Capitalize(item)
	}
	return result
}
