// Lissajous generates gif animations of random lissajous figure
package main

import (
	"os"

	"github.com/danbobs/go_training/lissajous"
)

func main() {
	lissajous.Create(os.Stdout)
}
