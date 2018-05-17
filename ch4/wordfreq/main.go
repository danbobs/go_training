// charcount counts the occurences of words
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int) // counts of words

	reader := bufio.NewReader(os.Stdin)
	scan := bufio.NewScanner(reader)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		counts[scan.Text()]++
	}

	fmt.Printf("word\t\tcount\n")
	for w, c := range counts {
		fmt.Printf("%s\t%d\n", w, c)
	}
}
