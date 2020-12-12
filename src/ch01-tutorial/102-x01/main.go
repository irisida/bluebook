package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// Exercise starts the loop from 0 to include Args[0]
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
