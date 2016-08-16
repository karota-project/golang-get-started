package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"ttsy", 28}

	// Marshal
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	log.Printf("%s", b)

	c, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		fmt.Errorf("%s", err)
	}

	log.Printf("%s", c)

	// Unmarshal
	var p2 Person
	err = json.Unmarshal([]byte(`{"Name":"ttsy", "Age":28}`), &p2)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	log.Println(p2)
}
