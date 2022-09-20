/*
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.
*/
package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter number between 1 to 100: ")
	var num int
	fmt.Scan(&num)

	list := make([]int, 0, 0)

	var count int
	for num > 0 {
		if num%2 == 1 {
			count++
		}
		if num/2 == 0 {
			list = append(list, num%2)
			num = 0
		} else {
			list = append(list, num%2)
			num = num / 2
		}
	}

	newArray := make([]int, 0, 0)
	newArray = append(newArray, 0)
	for i := len(list) - 1; i >= 0; i-- {
		newArray = append(newArray, list[i])
	}

	fmt.Println("Binary: ", newArray, "\nCount of 1's: ", count)

}
