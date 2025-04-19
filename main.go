package main

import (
	"flag"
	"log"
)

func main() {
	// usage: go run main.go --path ./your/path
	cmd := flag.String("path", "", "")
	flag.Parse()

	if string(*cmd) == "" {
		log.Fatal("No path was provided. Please use the '-path' flag in the go run main.go")
	}
}
