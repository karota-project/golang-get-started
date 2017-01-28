/*
ref : https://gist.github.com/Jxck/5237600
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// interface list
type Accessor interface {
	GetId() int
	SetId(interface{}) error
	GetName() string
	SetName(interface{}) error
}

// Struct
type Person struct {
	id   int
	name string
}

// Pointer type's method
func (p *Person) GetId() int {
	return p.id
}

func (p *Person) SetId(v interface{}) error {
	var err error

	switch v.(type) {
	case int:
		p.id = v.(int)
	case string:
		p.id, err = strconv.Atoi(v.(string))
	default:
		p.id = 0
		err = errors.New("Not supported.")
	}
	return err
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetName(v interface{}) error {
	var err error

	switch v.(type) {
	case int:
		p.name = strconv.Itoa(v.(int))
	case string:
		p.name = v.(string)
	default:
		p.name = ""
		err = errors.New("Not supported.")
	}
	return err
}

func main() {
	var person *Person = &Person{}
	person.SetName("kaiji")
	person.SetId(123)
	fmt.Println(person.GetName(), person.GetId())

	// use Accessor Interface
	var acsr Accessor = &Person{}
	err := acsr.SetName("Accessor")
	err = acsr.SetId("456")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(acsr.GetName(), acsr.GetId())
}
