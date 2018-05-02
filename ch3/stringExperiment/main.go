package main

import "fmt"

func main() {

	s := "ŃüŤş"
	fmt.Println(len(s))

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println(s[1:])

	r := []rune(s)
	fmt.Println(len(r))

}
