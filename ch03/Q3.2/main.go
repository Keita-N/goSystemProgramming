package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	random := rand.Reader
	file, err := os.Create("random.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.CopyN(file, random, 1024)
}
