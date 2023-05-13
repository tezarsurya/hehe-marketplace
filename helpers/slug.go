package helpers

import "strings"

func GenerateSlug(s string) string {
	lowercase := strings.ToLower(s)
	split := strings.Split(lowercase, " ")
	slug := strings.Join(split, "_")
	return slug
}
