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

		//symbol or comment
		if node.Type == TYPE_SYMBOL || node.Type == TYPE_COMMENT {
			var output strings.Builder
			traverserNodeBody(&output, &cursor, &node)
			bodyBuilder.WriteString(output.String())
			continue
		}

		// code
		if node.Type == TYPE_CODE {
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
				bodyBuilder.WriteString(val.Value)
			}
			bodyBuilder.WriteString("</" + element + ">")
			cursor += len(*node.Body) + 1

			continue
		}

		// string, space, newline
		bodyBuilder.WriteString(node.Value)
		cursor++

	}

	result := bodyBuilder.String()
	return &result
}

// ! core
func traverserNodeBody(output *strings.Builder, cursor *int, currentNode *ASTNode) {
	if currentNode.Type == TYPE_SYMBOL {
		DOMElement, ok := u.HTMLEquivalents[currentNode.Value]
		if !ok {
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

	for _, val := range *currentNode.Body {
		if len(*val.Body) > 0 {
			traverserNodeBody(output, cursor, &val)
		} else {
			// its just a string or empty character
			if val.Type != TYPE_COMMENT {
				output.WriteString(val.Value)
			}
		}
	}

	if currentNode.Type == TYPE_SYMBOL {
		DOMElement := u.HTMLEquivalents[currentNode.Value]
		output.WriteString(toggleHtmlSymbols(cursor, DOMElement, true))
	} else {
		output.WriteString(closeHtmlComments(cursor, currentNode.Value))
	}
}
