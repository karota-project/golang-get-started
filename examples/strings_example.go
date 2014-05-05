package main

import (
	"fmt"
	"strings"
)

func main() {
	/* func Contains(s, substr string) bool */
	fmt.Println(strings.Contains("abcdefg", "cd"))
	fmt.Println(strings.Contains("abcdefg", "ag"))

	/* func Index(s, sep string) int */
	fmt.Println(strings.Index("abcdefg", "e"))

	/* func Join(a []string, sep string) string */
	s := []string{"a", "b", "c"}
	fmt.Println(strings.Join(s, ""))

	/* func Replace(s, old, new string, n int) string */
	fmt.Println(strings.Replace("red red blue red", "red", "green", 2))

	/* func Split(s, sep string) []string */
	fmt.Printf("%s\n", strings.Split("a,b,c,d,e,f,g", ","))
	fmt.Printf("%s\n", strings.Split("abcdefg", ""))

	/* func Trim(s string, cutset string) string */
	fmt.Printf("%s\n", strings.Trim("abcdefg", "defg"))
}

/*
   $ go run string_example.go
   true
   false
   4
   abc
   green green blue red
   [a b c d e f g]
   [a b c d e f g]
   abc
*/
