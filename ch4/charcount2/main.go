// charcount counts the occurences of unicode characters
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {

	letterCounts := make(map[rune]int) // counts of letters
	digitCounts := make(map[rune]int)  // counts of digits
	otherCounts := make(map[rune]int)  // counts of all other char types

	var utfLen [utf8.UTFMax + 1]int // counts of lengths of utf8 encodings
	invalid := 0                    // counts of invalid characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("charcount: %v", err)
			os.Exit(0)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		switch {
		case unicode.IsLetter(r):
			letterCounts[r]++
		case unicode.IsDigit(r):
			digitCounts[r]++
		default:
			otherCounts[r]++
		}

		utfLen[n]++
	}

	printCharCount(letterCounts)
	printCharCount(digitCounts)
	printCharCount(otherCounts)

	fmt.Printf("rune\tcount\n")
	for i, n := range utfLen {
		fmt.Printf("%d\t%d\n", i, n)
	}

	if invalid > 0 {
		fmt.Printf("Invalid chars: %d", invalid)
	}
}

func printCharCount(counts map[rune]int) {
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

}
