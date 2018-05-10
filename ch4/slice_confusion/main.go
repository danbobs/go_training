package main

import (
	"fmt"
)

func main() {

	s := []int{1, 2, 3, 4, 5}
	s2 := s[:]
	s3 := s[0:2]

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)

	s[1] = 9

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)

}
