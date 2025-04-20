package utils

import (
	"log"
	"os"
	"path/filepath"
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
