package compiler

import (
	u "md_to_html/utils"
	"strings"
)

func TransformToHTMLCode(ASTree *[]ASTNode) *string {
	var bodyBuilder strings.Builder
	var cursor int

	for cursor < len(*ASTree) {
		node := (*ASTree)[cursor]

		if node.Type == "Symbol" || node.Type == "Comment" {
			var output strings.Builder
			traverserNodeBody(&output, &cursor, &node, ASTree)
			bodyBuilder.WriteString(output.String())
		} else {
			nextVal := getNextVal(ASTree, &cursor)
			resString := identStrings(&cursor, &node, nextVal)
			bodyBuilder.WriteString(resString)
			cursor++
		}
	}

	result := bodyBuilder.String()
	return &result
}

// ! core
func traverserNodeBody(output *strings.Builder, cursor *int, currentNode *ASTNode, ASTree *[]ASTNode) {
	// fmt.Println("CURRENT NODE: ", currentNode)
	if currentNode.Type == "Symbol" {
		DOMElement := u.HTMLEquivalents[currentNode.Value]
		output.WriteString(toggleHtmlSymbols(cursor, DOMElement, false))
	} else {
		// else, its a Comment
		output.WriteString(currentNode.Value)
		*cursor++
	}

	for idx, val := range *currentNode.Body {
		if len(*val.Body) > 0 {
			traverserNodeBody(output, cursor, &val, ASTree)

		} else {
			// its just a string or empty character
			if val.Type != "Comment" {
				nextVal := getNextVal(currentNode.Body, &idx)
				res := identStrings(cursor, &val, nextVal)
				output.WriteString(res)
			}
		}
	}

	if currentNode.Type == "Symbol" {
		DOMElement := u.HTMLEquivalents[currentNode.Value]
		output.WriteString(toggleHtmlSymbols(cursor, DOMElement, true))
	} else {
		output.WriteString(closeHtmlComments(cursor, currentNode.Value))

	}
}
