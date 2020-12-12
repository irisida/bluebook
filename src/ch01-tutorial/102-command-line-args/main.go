package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	/*
		start our loop from 1 on the basis that on the CLI
		Args[0] is the name of the script/process that is
		being run.
	*/
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
