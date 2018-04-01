// Dup: prints the text of each line that appears more than once in the standard input
// preceeded by the number of occurences
package main

import (
	"bufio"
	"fmt"
	"os"
)

var counts map[string]int

func main() {
	counts = make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		return
	}

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error loading file %s", filename)
		} else {
			countFileLines(file)
		}
		printDupes()
	}

}

func countFileLines(file *os.File) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func printDupes() {
	for text, count := range counts {
		if count > 1 {
			fmt.Fprintf(os.Stderr, "%s - %d\n", text, count)
		}
	}
}
