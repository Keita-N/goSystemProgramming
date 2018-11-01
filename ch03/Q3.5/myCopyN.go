package main

import (
	"io"
	"os"
	"strings"
)

func myCopyN(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
	return io.Copy(dst, io.LimitReader(src, n))
}

func main() {
	myCopyN(os.Stdout, strings.NewReader("MyCopyN Function Test"), 7)
}
