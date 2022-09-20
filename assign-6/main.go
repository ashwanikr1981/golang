/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].


*/
package main

import "fmt"

func getIndices(arr []int, target int) []int {
	newArray := make([]int, 0, 0)
	arr2 := arr
	for i1, v1 := range arr {
		for i2, v2 := range arr2 {
			if v1+v2 == target {
				if !contains(newArray, i1) {
					newArray = append(newArray, i1, i2)
				}
			}
		}
	}
	return newArray
}

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{2, 5, 7, 10, 15, 18}
	target := 9
	result := getIndices(arr, target)
	fmt.Println("Array: ", arr)
	fmt.Println("Target: ", target)
	fmt.Println("Result: ", result)
}
