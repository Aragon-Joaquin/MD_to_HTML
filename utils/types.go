package utils

type TypeOfData struct {
	Pattern []string
	// doesRepeat bool
}

var TypeOfSymbols = []string{"*", "_", "$", "=", "`",
	"~", "#", ">", "-",
	"[", "]", "(", ")",
	"/", "<", "!"}

var Symbols = map[string]TypeOfData{
	"*": {Pattern: []string{"*", "**", "***", "*/"}},
	"_": {Pattern: []string{"_", "__"}},
	"$": {Pattern: []string{"$"}},
	"=": {Pattern: []string{"=="}},
	"`": {Pattern: []string{"`", "```"}},
	"~": {Pattern: []string{"~~"}},
	"#": {Pattern: []string{"#", "##", "###", "####", "#####", "######"}},
	">": {Pattern: []string{">", "-->"}},
	"-": {Pattern: []string{"-", "---"}},
	"[": {Pattern: []string{"["}},
	"]": {Pattern: []string{"]"}},
	"(": {Pattern: []string{"("}},
	")": {Pattern: []string{")"}},
	"/": {Pattern: []string{"/*"}},
	"<": {Pattern: []string{"<!--"}},
}

var CommentCombined = map[string]int{
	"/*":   1,
	"*/":   1,
	"<!--": 1,
	"-->":  1,
}
