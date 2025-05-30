package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a [5]int

	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 300, 500}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	fmt.Println("2d: ", twoD[1][2])

	twoDArr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d: ss ", twoDArr)
	// fmt.Println("type of twodArr", typeof twoDArr[0])
	fmt.Println("type of  ", reflect.TypeOf(twoDArr))

	fmt.Println("type of  ", reflect.TypeOf(twoDArr[0]))

}
