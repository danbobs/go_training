package main

import (
	"fmt"
)

func main() {

	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v %v %v\n", s, len(s), cap(s))

	// len 1 capacity 1
	test1 := s[4:]
	fmt.Printf("%v %v %v\n", test1, len(test1), cap(test1))

	// len 0 cap 0 - not out of bounds!
	test2 := s[5:]
	fmt.Printf("%v %v %v\n", test2, len(test2), cap(test2))

	//test3 := s[6:] index out of bounds

}
