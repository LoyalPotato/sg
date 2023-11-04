package utils

import (
	"bytes"
	"io"
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

// Count number of newlines in reader
//
// Credits: https://stackoverflow.com/a/24563853/6251097
func LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
