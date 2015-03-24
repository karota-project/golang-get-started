package main

import (
	"fmt"
	"reflect"
)

func match(a, b interface{}) bool {
	aval := reflect.ValueOf(a)      // Value Data
	if aval.Kind() != reflect.Ptr { // Type Data
		return false
	}

	bval := reflect.ValueOf(b)
	if aval.Elem().Kind() != bval.Kind() {
		return false
	}

	return true
}

func main() {
	var i int
	var j float32
	var k string

	t1 := match(&i, 100)
	fmt.Println("t1 : ", t1)

	t2 := match(&j, 10.4)
	fmt.Println("t2 : ", t2)

	t3 := match(k, "abc")
	fmt.Println("t3 : ", t3)
}
