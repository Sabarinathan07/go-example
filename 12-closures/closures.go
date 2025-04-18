package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())

	// closures can be used to encapsulate state
	// and create private variables
	// this is a simple counter
	counter := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}
	c1 := counter()
	c2 := counter()
	fmt.Println(c1())
	fmt.Println(c2())
	fmt.Println(c1())
	fmt.Println(c2())

}
