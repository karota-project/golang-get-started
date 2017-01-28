package main

import (
	"log"
	"reflect"
)

func match(a, b interface{}) bool {
	log.Println("a : ", reflect.TypeOf(a))
	log.Println("b : ", reflect.TypeOf(b))

	av := reflect.ValueOf(a)
	if av.Kind() != reflect.Ptr {
		return false
	}

	bv := reflect.ValueOf(b)
	if av.Elem().Kind() != bv.Kind() {
		return false
	}

	return true
}

func main() {
	var i int
	var j float32
	var k string

	// *int and int
	im := match(&i, 100)
	log.Println("im : ", im)

	// *float32 and float64
	jm := match(&j, 10.4)
	log.Println("jm : ", jm)

	// *string and string
	km := match(k, "abc")
	log.Println("km : ", km)
}
