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
		} else if node.Type == "Code" {
			element := u.HTMLEquivalents[node.Value][0]
			var langAttr string = "bash"

			if len(*node.Body) >= 2 {
				lang := (*node.Body)[0]
				if lang.Value != "\n" && (*node.Body)[1].Value == "\n" {
					langAttr = lang.Value
					cursor++
					*node.Body = (*node.Body)[1:]
				}
			}
			bodyBuilder.WriteString("<" + element + " lang='" + langAttr + "'>")
			for _, val := range *node.Body {

				bodyBuilder.WriteString(val.Value + " ")
			}
			bodyBuilder.WriteString("</" + element + ">")
			cursor += len(*node.Body) + 1
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
	if currentNode.Type == "Symbol" {
		DOMElement := u.HTMLEquivalents[currentNode.Value]

		if len(DOMElement) == 0 {
			output.WriteString(currentNode.Value)
			*cursor++
			return
		}

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
