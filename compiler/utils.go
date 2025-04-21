package compiler

import (
	u "md_to_html/utils"
	"slices"
)

var lineSeparators = []string{"\r", "\n"}

func isString(char string) bool {
	if slices.Contains(u.TypeOfSymbols, char) {
		return false
	}

	isSeparator := slices.Contains(lineSeparators, char)

	if char == " " || isSeparator {
		return false
	}

	return true
}
