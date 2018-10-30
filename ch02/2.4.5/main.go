package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")

	// gzip
	file2, err2 := os.Create("test.txt.gz")
	if err2 != nil {
		panic(err2)
	}
	writer2 := gzip.NewWriter(file2)
	writer2.Header.Name = "test.txt"
	io.WriteString(writer2, "gzip.Writer example\n")
	writer2.Close()

	// buffer
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()

}
