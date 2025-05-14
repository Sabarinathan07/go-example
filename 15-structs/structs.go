package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(&person{name: "Fred", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age, "after")

	p := newPerson("Jon")
	fmt.Println(p.age)

	run2()
}

// run 
// step 1
//  go mod init go-example 
//  this is with root folder

// step 2
//  cd 15-structs
// go run .