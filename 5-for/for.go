package main

import (
	"fmt"
)

func main() {

	i:= 1
	for i <= 10 {
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()

	for j := 1; j <= 3; j++ {
		for k := 1; k <= 3; k++ {
			fmt.Print(j*k, " ")
		}
		fmt.Println()
	}

	for i:= range 5 {
		fmt.Println(i, " range")
	}

	fmt.Println()
	for {
		fmt.Println("infinite loop", i)

		if( i == 12) {
			break
		}
		i++;
	}
}