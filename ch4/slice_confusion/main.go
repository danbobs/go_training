package main

import (
	"fmt"
)

func main() {

	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v %v %v\n", s, len(s), cap(s))

	test1 := s[4:]
	fmt.Printf("%v %v %v\n", test1, len(test1), cap(test1))

	test2 := s[5:]
	fmt.Printf("%v %v %v\n", test2, len(test2), cap(test2))

}
