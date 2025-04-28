package main

import (
	"flag"
	"fmt"
	"log"
	c "md_to_html/compiler"
	u "md_to_html/utils"
	"os"
	"path/filepath"
)

// * usage: go run . --path "./your/path"
func main() {
	cmd := flag.String("path", "", "")
	flag.Parse()

	//! validate/verification steps
	// if _, err := os.Stat(c.FilePath); !errors.Is(err, os.ErrNotExist) {
	// 	log.Fatalln("file already exists, delete it or move it to another directory")
	// 	return
	// }

	var pathString string
	if string(*cmd) == "" {
		fmt.Fprintln(os.Stdout, u.Blue+`No path was provided. Please use the '-path' flag like this: <'go run . -path "./your/path"'>`)
		fmt.Fprintln(os.Stdout, u.Yellow+"Using fallback file 'example.md'..."+u.Reset)
		pathString = "./test/example.md"
	} else {
		fileExtension := filepath.Ext(*cmd)
		if fileExtension != ".md" && fileExtension != ".MD" {
			log.Fatalln("The file provided is not markdown: ", string(*cmd))
		}

		pathString = string(*cmd)
	}

	//! most important part of the program
	dataInfo := u.GetPath(pathString)
	tokens := c.TokenaizeAllLines(*dataInfo)
	ASTree := c.ParseToAST(*tokens)

	HTMLElements := c.TransformToHTMLCode(ASTree)
	pathname, err := c.CreateOutput(*HTMLElements)
	if err == nil {
		fmt.Fprintln(os.Stdout, u.Green+"File saved in: "+pathname+u.Reset)
	} else {
		fmt.Fprintln(os.Stdout, u.Red+err.Error()+u.Reset)
	}

}
