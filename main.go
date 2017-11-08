package main

import (
	"fmt"
	"io"
	"os"
)

type readFile struct {
	ptr     os.File
	content []byte
}

func main() {
	// assumes that we're typing a single filename after go run main.go at the command line
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
