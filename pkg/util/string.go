package util

import "strings"

func ConvertToSnakeCase(input string) string {
	var result strings.Builder
	for i, char := range input {
		if i > 0 && char >= 'A' && char <= 'Z' {
			result.WriteByte('_')
		}
		result.WriteByte(byte(char | 32)) // Convert to lowercase
	}
	return result.String()
}
