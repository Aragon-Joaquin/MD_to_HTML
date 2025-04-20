package main

import (
	"flag"
	"fmt"
	c "md_to_html/compiler"
	u "md_to_html/utils"
)

func main() {
	// usage: go run . --path "./your/path"
	cmd := flag.String("path", "", "")
	flag.Parse()

	var pathString string
	if string(*cmd) == "" {
		fmt.Println("No path was provided. Please use the '-path' flag in the go run main.go")
		pathString = "./test/example.md"
	} else {
		pathString = string(*cmd)
	}

	dataInfo := u.GetPath(pathString)
	tokens := c.TokenaizeAllLines(*dataInfo)
	c.ParseToAST(tokens)
}
