package main

import (
	"fmt"
)

type Person struct {
	id   int
	name string
}

func main() {
	/* int*/
	var i int = 1
	var ii = 2
	iii := 3
	fmt.Println(i, ii, iii)

	/* string */
	var s string = "a"
	var ss = "b"
	sss := "c"
	fmt.Println(s, ss, sss)

	/* int array */
	k := []int{1, 2, 3}
	fmt.Println(k)

	/* string array */
	strs := []string{"a", "b", "c"}
	fmt.Println(strs)

	/* structure */
	p := new(Person)
	p.id = 123
	p.name = "Kaiji"
	fmt.Println(p)

	/* format */
	fmt.Printf("%v,%v\n", i, s) // generic
	fmt.Printf("%s\n", s)       // string
	fmt.Printf("%f\n", 1.23)    // float
	fmt.Printf("%x\n", 0xff)    // hex

}
