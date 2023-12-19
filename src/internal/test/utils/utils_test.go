package utils_test

import (
	"strings"
	"testing"

	"github.com/LoyalPotato/sg/src/internal/utils"
)

func TestStringMatches(t *testing.T) {
	t.Run("string matches", func(t *testing.T) {
		s := "string"
		match := "string"
		matches := utils.StringMatches(s, match)
		if !matches {
			t.Fatalf(`StringMatches("string", "string") did not match`)
		}
	})

	t.Run("string matches with non graphic unicode", func(t *testing.T) {
		s := "string\t"
		match := "string"
		matches := utils.StringMatches(s, match)
		if !matches {
			t.Fatalf(`StringMatches("string\t", "string") did not match`)
		}
	})

	t.Run("string does not match", func(t *testing.T) {
		s := "string"
		match := "not string"
		matches := utils.StringMatches(s, match)
		if matches {
			t.Fatalf(`StringMatches("string", "not string") did not match`)
		}
	})
}

// TODO: Use a fuzzing test? Seems like it would be interesting for this type of function
// func FuzzStringMatches(f *testing.F) {
//
// }

func TestRemoveNonPrintables(t *testing.T) {
	t.Run("string has all printables", func(t *testing.T) {
		s := "string"
		out := utils.RemoveNonPrintables(s)
		if s != out {
			t.Fatalf(`RemoveNonPrintables("string") did not return the same output %q`, out)
		}
	})

	t.Run("string has some non printables", func(t *testing.T) {
		s := "string\t\n"
		out := utils.RemoveNonPrintables(s)
		want := "string"
		if out != want {
			t.Fatalf(`RemoveNonPrintables("string\t\n") did not remove all non printables to match %q`, want)
		}
	})
}

func TestLineCounter(t *testing.T) {
	t.Run("string has no lines", func(t *testing.T) {
		s := "no new line"
		sReader := strings.NewReader(s)
		lines, err := utils.LineCounter(sReader)
		if err != nil {
			t.Fatalf(`LineCounter error occurred: %v`, err)
		}

		if lines != 0 {
			t.Fatalf(`LineCounter with reader for %q didn't return 0 lines`, s)
		}
	})
}
