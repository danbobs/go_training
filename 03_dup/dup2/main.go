// Dup: prints the text of each line that appears more than once in the standard input
// preceeded by the number of occurences
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var counts map[string]int

func main() {
	counts = make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		return
	}

	for _, filename := range files {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v", err)
		} else {
			contents := string(bytes)
			countFileLines(contents)
		}
	}

	printDupes()
}

func countFileLines(contents string) {
	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		counts[line]++
	}
}

func printDupes() {
	for text, count := range counts {
		if count > 1 {
			fmt.Fprintf(os.Stderr, "%s - %d\n", text, count)
		}
	}
}
