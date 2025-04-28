package compiler

import (
	u "md_to_html/utils"
	"slices"
)

type ASTNode struct {
	ParentNode *ASTNode
	Type       string
	Value      string
	Body       *[]ASTNode
}

// ! this could be wrong
func ParseToAST(tokens []Token) *[]ASTNode {
	var cursor int = -1
	var recursiveToken func(parent *ASTNode, finalTree []ASTNode) []ASTNode

	recursiveToken = func(parent *ASTNode, finalTree []ASTNode) []ASTNode {
		cursor++

		if cursor >= len(tokens) {
			return finalTree
		}
		cuToken := tokens[cursor]

		if cuToken.Type == "String" {
			if parent != nil {
				*parent.Body = append(*parent.Body, createEmptyNode(cuToken, parent))
			} else {
				finalTree = append(finalTree, createEmptyNode(cuToken, nil))
			}
			return recursiveToken(parent, finalTree)
		}

		if cuToken.Type == "Symbol" {
			getSymbol := u.Symbols[cuToken.Value]
			var fullSymbol string
			alternatives := checkAllPosibilities(getSymbol.Pattern)

			// groups symbols like:
			// #####, _, <!--, etc...
			var counter int
			for idx, val := range alternatives {
				for range len(val) {
					cuToken := tokens[cursor+counter]
					counter++

					if slices.Contains(val, string(cuToken.Value)) {
						val = val[idx:]
						fullSymbol += cuToken.Value
					} else {
						break
					}
				}

				if slices.Contains(getSymbol.Pattern, fullSymbol) {
					break
				}

			}
			cursor += len(fullSymbol) - 1

			//opens a body for the symbol node
			if parent == nil {
				node := ASTNode{
					ParentNode: nil,
					Type:       isCommentType(fullSymbol),
					Value:      fullSymbol,
					Body:       &[]ASTNode{},
				}
				finalTree = append(finalTree, node)
				return recursiveToken(&node, finalTree)
			} else {
				// else, it closes it or reopens a new one

				parentPattern := u.Symbols[string(parent.Value[0])]

				if slices.Contains(parentPattern.Pattern, fullSymbol) {
					return recursiveToken(parentHasParent(parent), finalTree)
				} else {
					node := ASTNode{
						ParentNode: parent,
						Type:       isCommentType(fullSymbol),
						Value:      fullSymbol,
						Body:       &[]ASTNode{},
					}
					*parent.Body = append(*parent.Body, node)
					return recursiveToken(&node, finalTree)
				}

			}

		}

		if cuToken.Type == "NewLine" {
			finalTree = append(finalTree, createEmptyNode(cuToken, parent))
			return recursiveToken(parentHasParent(parent), finalTree)
		}

		return finalTree
	}

	ASTree := recursiveToken(nil, []ASTNode{})
	return &ASTree
}
