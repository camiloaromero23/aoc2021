package utils

import "regexp"

func RemoveDuplicateWhitespace(s string) string {
	spaceRegex := regexp.MustCompile(`\s+`)
	clean := spaceRegex.ReplaceAllString(s, " ")
	return clean
}
