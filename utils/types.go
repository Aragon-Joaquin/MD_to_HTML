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

	// complex symbols
	">": {Pattern: []string{">"}},
	"-": {Pattern: []string{"-", "---", "-->"}},
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

var HTMLEquivalents = map[string]([]string){
	"*":      {Italic},
	"**":     {Bold},
	"***":    {Italic, Bold},
	"_":      {Italic},
	"__":     {Italic, Bold},
	"$":      {Math},
	"==":     {Highlight},
	"`":      {Code},
	"~~":     {Deleted},
	"#":      {H1},
	"##":     {H2},
	"###":    {H3},
	"####":   {H4},
	"#####":  {H5},
	"######": {H6},
	">":      {Quote},
	"-":      {UnorderedList, List},
	"---":    {Break},
	"[":      {Anchor},
	//    "(": "a href", i'll think of this later on
}
