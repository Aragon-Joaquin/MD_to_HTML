package compiler

type TOKEN_TYPE string

const (
	TYPE_SYMBOL  TOKEN_TYPE = "Symbol"
	TYPE_NEWLINE TOKEN_TYPE = "NewLine"
	TYPE_STRING  TOKEN_TYPE = "String"
	TYPE_CODE    TOKEN_TYPE = "Code"
	TYPE_SPACE   TOKEN_TYPE = "Space"
	TYPE_COMMENT TOKEN_TYPE = "Comment"
)
