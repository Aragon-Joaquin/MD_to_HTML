package main

import (
	"flag"
	"fmt"
	c "md_to_html/compiler"
	u "md_to_html/utils"
	"os"
)

func main() {
	// usage: go run . --path "./your/path"
	cmd := flag.String("path", "", "")
	flag.Parse()

	var pathString string
	if string(*cmd) == "" {
		fmt.Fprintln(os.Stdout, u.Red+`No path was provided. Please use the '-path' flag like this: <'go run . -path "./your/path"'>`)
		fmt.Fprintln(os.Stdout, u.Yellow+"Using fallback file 'example.md'..."+u.Reset)
		pathString = "./test/example.md"
	} else {
		pathString = string(*cmd)
	}

	//! most important part of the program
	dataInfo := u.GetPath(pathString)
	tokens := c.TokenaizeAllLines(*dataInfo)
	ASTree := c.ParseToAST(tokens)
	c.TransformToHTMLCode(ASTree)
}
