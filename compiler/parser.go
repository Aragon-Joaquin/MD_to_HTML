package compiler

import (
	u "md_to_html/utils"
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

		// check if code
		if parent != nil && parent.Type == "Code" {
			if cursor+3 < len(tokens) {
				isBackticks := checkIfCode(tokens[cursor : cursor+3])
				if isBackticks {
					cursor += 3
					return recursiveToken(parentHasParent(parent), finalTree)
				}

				createString(parent, cuToken.Value, &finalTree)
				return recursiveToken(parent, finalTree)
			}

		}
		switch cuToken.Type {
		case "String":
			{
				createString(parent, cuToken.Value, &finalTree)
				return recursiveToken(parent, finalTree)
			}
		case "Symbol":
			{
				getSymbol := u.Symbols[cuToken.Value]

				var patternMatch string
				// var isBadSymbol bool
				var counter int
				for _, val := range getSymbol.Pattern {
					counter = 0
					// isBadSymbol = false
					patternMatch = ""

					for _, char := range val {
						if counter+cursor < len(tokens) {
							cuToken := tokens[counter+cursor]
							counter++

							if cuToken.Value == string(char) {
								patternMatch += cuToken.Value
							} else {
								break
							}
						} else {
							break
						}
					}

					// i have no idea how to check if the next symbol
					// makes sense for the context or not just to be taken as string
					if val == patternMatch || u.ClosesBy[patternMatch] != "" {
						break
					}
				}

				if patternMatch == "" {
					return recursiveToken(parent, finalTree)
				}
				cursor += len(patternMatch) - 1

				//check if its a bad symbol
				// if isBadSymbol{
				// 	createString(parent, patternMatch, &finalTree)
				// 	return recursiveToken(parent, finalTree)
				// }

				if parent != nil {
					closesBy := u.ClosesBy[parent.Value]

					if patternMatch == parent.Value || closesBy == patternMatch {
						return recursiveToken(parentHasParent(parent), finalTree)
					} else {
						node := ASTNode{
							ParentNode: parent,
							Type:       returnSymbolType(patternMatch),
							Value:      patternMatch,
							Body:       &[]ASTNode{},
						}
						*parent.Body = append(*parent.Body, node)
						return recursiveToken(&node, finalTree)
					}
				} else {
					node := ASTNode{
						ParentNode: nil,
						Type:       returnSymbolType(patternMatch),
						Value:      patternMatch,
						Body:       &[]ASTNode{},
					}
					finalTree = append(finalTree, node)
					return recursiveToken(&node, finalTree)
				}
			}

		default:
			{
				// most likely a newLine
				node := ASTNode{
					ParentNode: nil,
					Type:       cuToken.Type,
					Value:      cuToken.Value,
					Body:       &[]ASTNode{},
				}

				finalTree = append(finalTree, node)
				return recursiveToken(nil, finalTree)
			}
		}

	}

	ASTree := recursiveToken(nil, []ASTNode{})
	return &ASTree
}
