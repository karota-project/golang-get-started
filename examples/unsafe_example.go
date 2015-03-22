package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type T struct {
	value int
}

func Swap(dest **T, old, new *T) bool {
	udest := (*unsafe.Pointer)(unsafe.Pointer(dest))
	return atomic.CompareAndSwapPointer(udest,
		unsafe.Pointer(old),
		unsafe.Pointer(new),
	)
}

func main() {
	str := "hello, world."
	b := *(*[]byte)(unsafe.Pointer(&str))
	fmt.Println(b) // => [104 101 108 108 111 44 32 119 111 114 108 100 46]

	x := &T{42}
	n := &T{50}
	fmt.Println(*x, *n) // => {42} {50}

	p := x
	Swap(&x, p, n)
	fmt.Println(*x, *n) // => {50} {50}
}
