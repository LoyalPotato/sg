package utils

import (
	"strings"
	"unicode"
)

func StringMatches(s string, match string) bool {
	newStr := RemoveNonPrintables(s)

	return newStr == match
}

func RemoveNonPrintables(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, s)
}
