package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func foo(n int) {
	fmt.Println(n)
}

func main() {
	/* get file*/
	_, file, _, _ := runtime.Caller(1)
	fmt.Println(file)

	/* get prev func*/
	prev := reflect.ValueOf(foo)
	fmt.Println(runtime.FuncForPC(prev.Pointer()).Name())
}
