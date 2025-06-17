package compiler

import (
	u "md_to_html/utils"
	"slices"
)

type Token struct {
	Type  TOKEN_TYPE
	Value string
}

func TokenaizeAllLines(fileBytes []byte) *[]Token {
	var sliceOfTokens []Token
	var cursor int = 0

	for cursor < len(fileBytes) {
		char := string(fileBytes[cursor])

		if CheckForTypeSpace(char) {
			sliceOfTokens = append(sliceOfTokens, Token{
				Type:  TYPE_SPACE,
				Value: char,
			})
			cursor++
			continue
		}

		if slices.Contains(u.TypeOfSymbols, char) {
			sliceOfTokens = append(sliceOfTokens, Token{
				Type:  TYPE_SYMBOL,
				Value: char,
			})
			cursor++
			continue
		}

		//! check for line feed. (depending on the OS could change)
		//! "\r" is carriage return, and "\n" is line feed.

		if slices.Contains(lineSeparators, char) {
			sliceOfTokens = append(sliceOfTokens, Token{
				Type:  TYPE_NEWLINE,
				Value: "\n",
			})
			cursor++
			continue
		}

		var word string
		for {
			char := string(fileBytes[cursor])

			if isString(char) {
				word += char
				cursor++
				continue
			} else {
				sliceOfTokens = append(sliceOfTokens, Token{
					Type:  TYPE_STRING,
					Value: word,
				})
			}
			break

		}
	}

	return &sliceOfTokens

}
