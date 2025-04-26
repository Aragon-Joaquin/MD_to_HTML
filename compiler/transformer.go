package compiler

import (
	u "md_to_html/utils"
	"strings"
)

// ! probably i need to make this as a recursive function
func TransformToHTMLCode(ASTree *[]ASTNode) {
	var HTMLOutput string

	for _, val := range *ASTree {

		DOMElement := u.HTMLEquivalents[val.Value]
		if len(DOMElement) > 0 {
			var body strings.Builder

			for _, val := range DOMElement {
				body.WriteString("<" + val + ">")
			}
			body.WriteString("\n")

			for range *val.Body {

			}
		} else {
			HTMLOutput += val.Value
		}
	}

}
