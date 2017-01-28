package main

import (
	"fmt"
)

func main() {
	strs := []string{"a", "b", "c", "d", "e"}
    fmt.Println(strs[2:4])

    var slice []string
	for _, s := range strs {
		if s == "a" || s == "e" {
			slice = append(slice, s)
		}
	}
	fmt.Println(slice)

    var sliceM []int = make([]int, 3, 3)
    fmt.Println(len(sliceM), sliceM)
}
