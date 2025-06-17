package compiler

import (
	u "md_to_html/utils"
	"slices"
	"strings"
	"unicode/utf8"
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

// TRUE == is space or tabulation
//
// FALSE == is NewLine, Symbol, etc...
func CheckForTypeSpace(char string) bool {
	//! U+0009 is tabs &  U+0020 is space
	if r, _ := utf8.DecodeRuneInString(char); r == '\t' || r == ' ' {
		return true
	}
	return false
}

/*
* ------------------------
* 	PARSER FUNCTIONS
* ------------------------
 */

func returnSymbolType(pattern string) TOKEN_TYPE {
	if u.CommentCombined[pattern] > 0 {
		return TYPE_COMMENT
	}

	if pattern == "```" {
		return TYPE_CODE
	}

	return TYPE_SYMBOL
}

func parentHasParent(node *ASTNode) *ASTNode {

	if node == nil || node.ParentNode == nil {
		return nil
	}

	if node.ParentNode.ParentNode == nil {
		return nil
	}

	return node.ParentNode.ParentNode
}

func createString(parent *ASTNode, Value string, finalTree *[]ASTNode) *ASTNode {
	node := ASTNode{
		ParentNode: parent,
		Type:       TYPE_STRING,
		Value:      Value,
		Body:       &[]ASTNode{},
	}

	if parent != nil {
		*parent.Body = append(*parent.Body, node)
	} else {
		*finalTree = append(*finalTree, node)
	}

	return &node
}

func checkIfCode(nodes []Token) bool {
	var isBackticks bool = true

	for _, val := range nodes {
		if val.Value != "`" {
			isBackticks = false
			break
		}
	}

	return isBackticks
}

/*
* ------------------------
* 	TRANSFORMER FUNCTIONS
* ------------------------
 */

func toggleHtmlSymbols(cursor *int, DOMElement []string, closeElements bool) string {
	var bodyBuilder strings.Builder

	for _, val := range DOMElement {
		if closeElements {
			bodyBuilder.WriteString("</" + val + ">")
			bodyBuilder.WriteString("\n")
		} else {
			bodyBuilder.WriteString("<" + val + ">")
		}
		*cursor++
	}
	return bodyBuilder.String()
}

func closeHtmlComments(cursor *int, DOMElement string) string {
	var bodyBuilder strings.Builder

	commentType := u.ClosesBy[DOMElement]

	bodyBuilder.WriteString(commentType)
	*cursor++

	return bodyBuilder.String()
}
