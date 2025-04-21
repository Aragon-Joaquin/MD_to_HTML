package compiler

import (
	"fmt"
	u "md_to_html/utils"
	"slices"
	"strings"
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
				finalTree = append(finalTree, createEmptyNode(cuToken, parent))
			}
			return recursiveToken(parent, finalTree)
		}

		if cuToken.Type == "Symbol" {
			getSymbol := u.Symbols[cuToken.Value]
			var fullSymbol string
			alternatives := checkAllPosibilities(getSymbol.Pattern)

			//! still working on this
			for idx, val := range alternatives {
				cuToken := tokens[cursor]
				cursor++

				fmt.Println(string(fullSymbol + cuToken.Value))
				if slices.Contains(getSymbol.Pattern, string(fullSymbol+cuToken.Value)) {
					fullSymbol += cuToken.Value
				} else {
					break
				}
			}

			fmt.Printf("fullSymbol: %v\n", fullSymbol)

			if parent == nil {
				node := createEmptyNode(cuToken, parent)
				finalTree = append(finalTree, node)
				return recursiveToken(&node, finalTree)
			} else {
				matchValue := strings.Split(parent.Value, "")
				var patternMatch string

				for idx := range matchValue {
					cuToken := tokens[cursor]
					cursor++

					if matchValue[idx] == cuToken.Value {
						patternMatch += cuToken.Value
						continue
					} else {
						break
					}
				}

				if patternMatch == parent.Value {
					return recursiveToken(parent.ParentNode, finalTree)
				} else {
					node := createEmptyNode(cuToken, parent)
					*parent.Body = append(*parent.Body, node)
					return recursiveToken(parent, finalTree)
				}

			}

		}

		if cuToken.Type == "NewLine" {
			finalTree = append(finalTree, createEmptyNode(cuToken, parent))
			return recursiveToken(nil, finalTree)
		}

		return finalTree
	}

	ASTree := recursiveToken(nil, []ASTNode{})
	return &ASTree
}

func createEmptyNode(cuToken Token, parent *ASTNode) ASTNode {
	return ASTNode{
		ParentNode: parent,
		Type:       cuToken.Type,
		Value:      cuToken.Value,
		Body:       &[]ASTNode{},
	}
}

/*
! "why does this function even exists?"

Some characters like "<" have a pattern which contains multiples characters.
And some have more than one pattern.
*/
func checkAllPosibilities(pattern []string) [][]string {
	var alternatives [][]string

	for idx := range pattern {
		test := strings.Split(pattern[idx], "")
		alternatives = append(alternatives, test)
	}

	return alternatives
}
