/*
ref : https://gist.github.com/Jxck/5237600
*/

package main

import "fmt"

type Accessor interface {
	GetId() int
	SetId(int)
	GetName() string
	SetName(string)
}

// Struct
type Person struct {
	id   int
	name string
}

// Method
func (p *Person) GetId() int {
	return p.id
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetId(id int) {
	p.id = id
}

func (p *Person) SetName(name string) {
	p.name = name
}

func main() {
	// ポインタを使う
	var person *Person = &Person{}
	person.SetName("kaiji")
	person.SetId(123)
	fmt.Println(person.GetName(), person.GetId())

	// use Accessor Interface
	// Accessor型でメソッドリストにアクセスできる
	var acsr Accessor = &Person{}
	acsr.SetName("Accessor")
	acsr.SetId(456)
	fmt.Println(acsr.GetName(), acsr.GetId())
}
