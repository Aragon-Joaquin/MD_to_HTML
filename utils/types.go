package utils

//! there's no use for the int part. I just wanted to avoid slices just to make it more
var Symbols = map[string]int{
	"*": 1,
	"_": 1,
	"$": 1,
	"=": 1,
	"`": 1,
	"~": 1,
	"#": 1,
	">": 1,
	"-": 1,
	"[": 1,
	"]": 1,
	"(": 1,
	")": 1,

	//comment symbols:
	"/": 1,
	"<": 1,
	"!": 1,
}

var CommentCombined = map[string]int{
	"/*":   1,
	"*/":   1,
	"<!--": 1,
	"-->":  1,
}
