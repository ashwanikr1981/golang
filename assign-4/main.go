/*
Generate the Series 1A2B3C4D5E6F.......
using  goroutine and channel  .
Routine 1 shall generate A,B C D E F and
routine 2 shall generate 1 ,2,3,4,5,6.......
Both this output join together in the output
*/
package main

import (
	"fmt"
)

func printNumber(num chan int) {

	for i := 1; i <= 6; i++ {
		num <- i
	}
	close(num)

}

func printAlphabets(alpha chan string) {

	for i := 'A'; i <= 'F'; i++ {
		alpha <- string(i)
	}
	close(alpha)

}

func main() {

	num := make(chan int)

	alpha := make(chan string)

	go printNumber(num)

	go printAlphabets(alpha)

	for {
		v, ok := <-num
		if !ok {
			return
		}
		fmt.Print(v)

		v1, ok1 := <-alpha
		if !ok1 {
			return
		}
		fmt.Print(v1)
	}
}
