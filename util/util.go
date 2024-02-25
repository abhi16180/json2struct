package util

import (
	"strings"
)

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
