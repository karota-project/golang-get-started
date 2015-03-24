package main

/*
#include <stdio.h>
int squared(int n) {
	return n * n;
};
*/
import "C"
import "fmt"

func main() {
	r := C.squared(5)
	fmt.Println(r)
}
