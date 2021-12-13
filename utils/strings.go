package utils

import (
	"regexp"
	"sort"
)

func RemoveDuplicateWhitespace(s string) string {
	spaceRegex := regexp.MustCompile(`\s+`)
	clean := spaceRegex.ReplaceAllString(s, " ")
	return clean
}

func SortStringByCharacter(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
