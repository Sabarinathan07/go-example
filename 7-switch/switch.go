package main

import (
	"fmt"
	"time"
)

func main() {
	// Switch statement
	var num int = 2
	switch num {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
	default:
		fmt.Println("Number is not 1, 2, or 3")
	}

	// Switch with no condition
	switch {
	case num < 0:
		fmt.Println("Number is negative")
	case num == 0:
		fmt.Println("Number is zero")
	default:
		fmt.Println("Number is positive")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
