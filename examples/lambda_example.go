package main

import "fmt"

func main() {

	y := 3
	func(x int) {
		fmt.Println(x + y) // => 5
	}(2)

	l := func(y int) int {
		return (y * y)
	}
	fmt.Println(l(y)) // => 9
}
