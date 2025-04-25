package main

import (
	"fmt"
)

// type person struct {
// 	name string
// 	age  int
// }

func birthday(p *person) {
	p.age += 1
	fmt.Println(p.name, "just turned", p.age)
}

func run2() {
	// Create a pointer to a person using &person{...}
	p := &person{name: "Alice", age: 29}

	// Print the original details
	fmt.Println("Before birthday:", p.name, p.age)

	// Pass the pointer to a function to modify the original struct
	birthday(p)

	// Print again to confirm change
	fmt.Println("After birthday:", p.name, p.age)
}
