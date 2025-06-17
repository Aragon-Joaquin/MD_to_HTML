package compiler

import (
	"fmt"
	u "md_to_html/utils"
)

type ASTNode struct {
	ParentNode *ASTNode
	Type       TOKEN_TYPE
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
		if parent != nil && parent.Type == TYPE_CODE {
			if cursor+3 < len(tokens) {
				fmt.Println(tokens[cursor : cursor+3])
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
		case TYPE_STRING:
			{
				createString(parent, cuToken.Value, &finalTree)
				return recursiveToken(parent, finalTree)
			}
		case TYPE_SYMBOL:
			{
				getSymbol := u.Symbols[cuToken.Value]

				var patternMatch string
				var counter int
				for _, val := range getSymbol.Pattern {
					counter = 0
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
					if val == patternMatch || u.ClosesBy[patternMatch] != "" {
						break
					}
				}

				if patternMatch == "" {
					return recursiveToken(parent, finalTree)
				}

				cursor += len(patternMatch) - 1

				fmt.Println("MATCHING: ", patternMatch)

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

		case TYPE_SPACE:
			{
				node := ASTNode{
					ParentNode: parent,
					Type:       cuToken.Type,
					Value:      cuToken.Value,
					Body:       &[]ASTNode{},
				}

				if parent != nil {
					*parent.Body = append(*parent.Body, node)
				} else {
					finalTree = append(finalTree, node)
				}
				return recursiveToken(parent, finalTree)
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
