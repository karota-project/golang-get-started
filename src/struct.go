package main

import (
	"log"
)

type Foo struct {
	x int
}

type Bar struct {
	Foo
	y int
}

func main() {
	bar := &Bar{Foo{1}, 2}

	log.Println(bar.Foo.x, bar.y)

	// embedded struct
	log.Println(bar.x, bar.y)
}
