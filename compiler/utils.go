package compiler

import (
	u "md_to_html/utils"
	"slices"
	"strings"
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

/*
* ------------------------
* 	PARSER FUNCTIONS
* ------------------------
 */

func returnSymbolType(pattern string) string {
	if u.CommentCombined[pattern] > 0 {
		return "Comment"
	}

	if pattern == "```" {
		return "Code"
	}

	return "Symbol"
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
		Type:       "String",
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

func getNextVal(slice *[]ASTNode, cursor *int) *ASTNode {
	if cursor == nil || *cursor+1 >= len(*slice) {
		return nil
	}

	nextVal := (*slice)[*cursor+1]
	return &nextVal
}

func identStrings(cursor *int, node *ASTNode, nextVal *ASTNode) string {
	var output strings.Builder
	output.WriteString(node.Value)

	if nextVal == nil {
		return output.String()
	} else {
		output.WriteString(" ")
		return output.String()
	}
}
