package main
import (
	"fmt"
}

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1.0 + 1.0 =", 1.0 + 1.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Println(1 << 1) // 2
	fmt.Println(1 << 2) // 4
}
