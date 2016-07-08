package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var rw ReadWriter

	// os.Stdout 实现了 Writer
	rw = os.Stdout

	fmt.Fprintf(rw, "hello, writer\n")
}
