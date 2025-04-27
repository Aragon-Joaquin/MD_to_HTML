package compiler

import (
	"bufio"
	"os"
)

const (
	FileDir         = "./output/"
	FileName string = "HTMLResult.html"
	FilePath string = FileDir + FileName
)

func CreateOutput(elements string) (string, error) {
	f, err := os.Create(FilePath)
	if err != nil {
		return "", err
	}

	w := bufio.NewWriter(f)
	defer w.Flush()
	_, err = w.WriteString(`<!DOCTYPE html>
	<!-- File automatically parsed -->
		<main>
		` + elements + `
		</main>`)

	if err != nil {
		return "", err
	}

	return f.Name(), nil
}
