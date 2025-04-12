package main

import (
	"fmt"
)
func main() {
	// Variables
	var a string = "Hello"
	var b int = 42
	var c float64 = 3.14
	var d bool = true
	var e rune = 'A'
	var f complex128 = 1 + 2i

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	a = "Hello, World!"
	fmt.Println(a + " modified")


	// Short variable declaration
	g := "World"
	h := 100
	i := 2.71828
	j := false
	k := 'B'
	l := 3 + 4i

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	// Multiple variable declaration
	var m, n, o int = 1, 2, 3
	p, q, r := "Go", "is", "awesome"

	fmt.Println(m, n, o)
	fmt.Println(p, q, r)
}