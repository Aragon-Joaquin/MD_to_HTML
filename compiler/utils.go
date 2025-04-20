package compiler

import (
	u "md_to_html/utils"
	"slices"
)

var lineSeparators = []string{"\r", "\n"}

func isString(char string) bool {
	if u.Symbols[char] > 0 {
		return false
	}

	isSeparator := slices.Contains(lineSeparators, char)

	if char == " " || isSeparator {
		return false
	}

	return true
}
