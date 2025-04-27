package compiler

import (
	u "md_to_html/utils"
	"strings"
)

func TransformToHTMLCode(ASTree *[]ASTNode) *string {
	var createHTMLElements func() strings.Builder
	var bodyBuilder strings.Builder
	var cursor int

	createHTMLElements = func() strings.Builder {
		for cursor < len(*ASTree) {
			node := (*ASTree)[cursor]

			if node.Type == "Symbol" || node.Type == "Comment" {
				var output strings.Builder
				if node.Type == "Symbol" {
					DOMElement := u.HTMLEquivalents[node.Value]
					//process elements like <b>, <i>
					bodyBuilder.WriteString(toggleHtmlSymbols(&cursor, DOMElement, false))

					for range *node.Body {
						res := createHTMLElements()
						output.WriteString(res.String())
					}

					bodyBuilder.WriteString(toggleHtmlSymbols(&cursor, DOMElement, true))

				} else {
					bodyBuilder.WriteString(node.Value)
					cursor++

					for range *node.Body {
						res := createHTMLElements()
						output.WriteString(res.String())
					}

					bodyBuilder.WriteString(closeHtmlComments(&cursor, node.Value))
				}

			} else {
				bodyBuilder.WriteString(node.Value)
				cursor++
			}
		}

		return bodyBuilder
	}

	createHTMLElements()
	result := bodyBuilder.String()
	return &result
}
