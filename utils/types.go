package utils

type TypeOfData struct {
	Pattern []string
}

var TypeOfSymbols = []string{"*", "_", "$", "=", "`",
	"~", "#", ">", "-",
	"[", "]", "(", ")",
	"/", "<", "!"}

var Symbols = map[string]TypeOfData{

	//get closes by the same pattern
	"*": {Pattern: []string{"*/", "***", "**", "*"}},
	"_": {Pattern: []string{"__", "_"}},
	"$": {Pattern: []string{"$"}},
	"=": {Pattern: []string{"=="}},
	"`": {Pattern: []string{"```", "`"}},
	"~": {Pattern: []string{"~~"}},
	"#": {Pattern: []string{"######", "#####", "####", "###", "##", "#"}},

	// wont repeat at the end
	">": {Pattern: []string{">"}},
	"-": {Pattern: []string{"-->", "---", "-"}},
	"[": {Pattern: []string{"["}},
	"]": {Pattern: []string{"]"}},
	"(": {Pattern: []string{"("}},
	")": {Pattern: []string{")"}},
	"/": {Pattern: []string{"/*"}},
	"<": {Pattern: []string{"<!--"}},
}

var ClosesBy = map[string]string{
	"/*":   "*/",
	"<!--": "-->",
	"(":    ")",
	"[":    "]",
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
	"`":      {KeyboardKey},
	"```":    {Code},
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
