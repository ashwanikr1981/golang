/*
	Given two integers a and b, return the sum of the two integers without using the operators + and -.
*/

package main

import (
	"fmt"
)

func add(x int, y int) int {
	fmt.Println(x, y)

	for i := 1; i <= y; i++ {
		x++
	}

	return x
}

func main() {

	a := 5
	b := 10
	c := add(a, b)
	fmt.Println("Adding of ", a, " and ", b, " is ", c)
}
