package main

import (
	"fmt"
)

func main() {
	/* map structure */
	m := map[string]string{"name": "karota", "established": "2013","location": "tokyo"}

	for k, v := range m {
		fmt.Println("key : ", k,", value : ",  v)
	}

	if _, ok := m["established"]; ok {
		fmt.Println(ok)
	}

}
