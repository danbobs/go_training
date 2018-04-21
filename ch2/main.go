// testing popcount
package main

import (
	"fmt"

	"github.com/danbobs/go_training/ch2/popcount"
)

func main() {
	x := uint64(244)
	fmt.Printf("Popcount of %v (%b) is %d", x, x, popcount.PopCount3(x))
}
