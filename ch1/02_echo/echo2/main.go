// Echo2 prints it's command line arguments, different variant
package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	sep := " "
	for _, arg := range os.Args[1:] {
		s += fmt.Sprintf("%s%s", arg, sep)
	}

	fmt.Println(s)
}
