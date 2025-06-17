package main

import (
	"errors"
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
	if _, err := os.Stat(c.FilePath); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			os.Remove(c.FilePath)
			fmt.Fprintln(os.Stdout, u.Red+"Previous file deleted."+u.Reset)
		} else {
			log.Fatalln(err.Error())
			return
		}
	}

	var pathString string
	if string(*cmd) == "" {
		fmt.Fprintln(os.Stdout, u.Blue+`No path was provided. Please use the '-path' flag like this: <$ go run . -path "./your/path">`)
		fmt.Fprintln(os.Stdout, u.Yellow+"Using fallback file 'example.md'..."+u.Reset)
		pathString = "./test/example.md"
	} else {
		fileExtension := filepath.Ext(*cmd)
		if fileExtension != ".md" && fileExtension != ".MD" {
			log.Fatalln("The file provided is not markdown: ", string(*cmd))
		}

		pathString = string(*cmd)
	}

	dataInfo := u.GetPath(pathString)

	//! compiler steps - most important part of the program
	tokens := c.TokenaizeAllLines(*dataInfo)
	ASTree := c.ParseToAST(*tokens)

	// my debug tool :clueless:
	// for _, val := range *ASTree {
	// 	if len(*val.Body) > 0 {
	// 		fmt.Fprintln(os.Stdout, u.Green+val.Value+" :", val.Type+u.Reset)
	// 		for _, lol := range *val.Body {
	// 			fmt.Println(string(lol.Type) + " " + lol.Value)
	// 		}
	// 		fmt.Println("")
	// 	}
	// }

	HTMLElements := c.TransformToHTMLCode(ASTree)

	//! output the file
	pathname, err := c.CreateOutput(*HTMLElements)
	if err == nil {
		fmt.Fprintln(os.Stdout, u.Green+"File saved in: "+pathname+u.Reset)
	} else {
		fmt.Fprintln(os.Stdout, u.Red+err.Error()+u.Reset)
	}

}
