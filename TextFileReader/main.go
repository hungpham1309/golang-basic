package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args)

	file, err := os.Open(os.Args[1])

	if err != nil {
		return
	}

	io.Copy(os.Stdout, file)
}
