package compiler

import (
	"fmt"
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

func isCommentType(pattern string) string {
	if u.CommentCombined[pattern] > 0 {
		return "Comment"
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
	fmt.Println("NEXT VAL:", nextVal)
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
