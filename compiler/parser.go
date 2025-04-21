package compiler

// u "md_to_html/utils"

type ASTNode struct {
	Type  string
	Value string
	Body  *[]ASTNode
}

func ParseToAST(tokens []Token) {
	var cursor int = 0
	var ASTree []ASTNode

	recursiveToken("**", []ASTNode{})

	// for cursor < len(tokens) {
	// 	currentToken := tokens[cursor]
	// 	token := u.Symbols[currentToken.Type].Pattern

	// }
}

//! testing logic, there's obviously gonna be errors
var cursor int = 0
var ASTree []ASTNode
var tokens []Token

func recursiveToken(parent string, word []ASTNode) []ASTNode {

	if (cursor + 1) > len(tokens) {
		return word
	}

	nextToken := tokens[cursor+1]

	//! other two possibilities are NewLine or String
	if nextToken.Type == "String" {
		word = append(word, ASTNode{
			Value: nextToken.Value,
			Type:  nextToken.Value,
			Body:  &[]ASTNode{},
		})
		recursiveToken(parent, word)
	}

	if nextToken.Type == "Symbol" {

	}

	return word
}
