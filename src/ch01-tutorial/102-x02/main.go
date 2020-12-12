package main

import (
	"fmt"
	"os"
)

func main() {
	// Exercise starts the loop from 0 to include Args[0]
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%d - %s \n", i, os.Args[i])
	}
}
