package utils

import "strings"

func CamelCase(word string) string {
	return strings.Title(strings.ToLower(word))
}
