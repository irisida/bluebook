package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "" {
			// add a controller to be able to quite the infinite loop
			// if the input is a bank entry the loop ends.
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		// take only entries with a count more than 1, ie, duplicates.
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
