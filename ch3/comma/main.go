// Comma inserts "thousands" commas into correct position of numeric string
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	count := 0

	runes := []rune(s)
	if len(s) <= 3 {
		return s
	}

	its := len(s)/3 + 1
	sectionLen := len(s) % 3
	for i := 0; i < its; i++ {
		sub := runes[count : count+sectionLen]
		count += sectionLen
		sectionLen = 3
		subStr := string(sub)
		buf.WriteString(subStr)
		if i != its-1 {
			buf.WriteString(",")
		}
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345678"))

}
