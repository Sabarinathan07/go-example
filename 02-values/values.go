package main

import (
	"fmt"
)

func main() {
	fmt.Println("go " + "lang")
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1.0 + 1.0 =", 1.0 + 1.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Println(1 << 1) // 2 0001 << 1 = 0010
	fmt.Println(1 << 2) // 4 0001 << 2 = 0100
	fmt.Println(1 << 3) // 8 0001 << 3 = 1000
	fmt.Println(1 << 4) // 16 0001 << 4 = 0001 0000
	fmt.Println(1 << 5) // 32 0001 << 5 = 0010 0000
	fmt.Println(1 << 6) // 64 0001 << 6 = 0100 0000
	fmt.Println(1 << 7) // 128 0001 << 7 = 1000 0000	
	fmt.Println(1 << 8) // 256 0001 << 8 = 0001 0000 0000
}
