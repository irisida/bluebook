package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int, fileList map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		counts[line]++

		// Add the filename to the list of files
		// a case was identified in if it has
		// not been added already.
		if !fileFoundInFileList(f.Name(), fileList[line]) {
			fileList[line] = append(fileList[line], f.Name())
		}
	}
}

// fileFoundInFileList - checks if file is in the file list
func fileFoundInFileList(target string, possibilities []string) bool {
	for _, t := range possibilities {
		if target == t {
			return true
		}
	}
	return false
}

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	fileList := make(map[string][]string)

	// if there are no files.
	// We will use os.Stdin
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileList) // fileList will default to [/dev/stdin] in this case
	} else {
		// Files have been passed to the program,
		// we will *not* use os.Stdin
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileList)
			f.Close()
		}
	}

	// print the duplicates
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, fileList[line])
		}
	}
}
