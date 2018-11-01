package main

import (
	"archive/zip"
	"os"
)

func main() {
	file, err := os.Create("test.txt.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	writer.Write([]byte("Hello."))
}
