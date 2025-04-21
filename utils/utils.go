package utils

import (
	"log"
	"os"
	"path/filepath"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func GetPath(path string) *[]byte {
	exPath, err := filepath.Abs(path)

	if err != nil {
		log.Fatalln("Fatal error ocurred: ", err)
	}

	fileByte, err := os.ReadFile(exPath)

	if err != nil {
		log.Fatalln("Fatal error ocurred: ", err)
	}

	return &fileByte
}

// html elements
const (
	Italic        = "i"
	Bold          = "b"
	Math          = "math"
	Highlight     = "mark"
	Code          = "code"
	Deleted       = "del"
	H1            = "h1"
	H2            = "h2"
	H3            = "h3"
	H4            = "h4"
	H5            = "h5"
	H6            = "h6"
	Quote         = "blockquote"
	UnorderedList = "ul"
	List          = "li"
	Anchor        = "a"
	Break         = "hr"
)
