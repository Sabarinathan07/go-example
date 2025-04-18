package main

import (
	"fmt"
	"maps"
)

func main() {
	// var m map[string]int
	// fmt.Println("emp:", m)
	// fmt.Println("len:", len(m))
	// fmt.Println("get:", m["key"])
	// fmt.Println("set:", m["key"] = 100)
	// fmt.Println("get:", m["key"])
	// fmt.Println("len:", len(m))

	m := make(map[string]int)
	fmt.Println("emp:", m)
	fmt.Println("len:", len(m))
	fmt.Println("get:", m["key"])
	m["key"] = 100
	fmt.Println("set:", m)
	fmt.Println("get:", m["key"])
	fmt.Println("len:", len(m))

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("dcl:", m2)

	m3 := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	fmt.Println("dcl:", m3)

	if maps.Equal(m2, m3) {
		fmt.Println("m2 == m3")
	} else {
		fmt.Println("m2 != m3")
	}

	delete(m3, "foo")
	fmt.Println("del:", m3)

	fmt.Println("foo:m3 ", m3["foo"])
	valueOfM3, ok := m3["foo"]
	if ok {
		fmt.Println(valueOfM3, ok)
	} else {
		fmt.Println("not found")
	}

	for k, v := range m3 {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
