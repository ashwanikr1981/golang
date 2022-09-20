/*
Create Two goroutine ....
Routine 1 should generate the Series from 1-100 and routine 2 check whether
its numbers are prime number or not and get it print
*/

package main

import (
	"fmt"
	"time"
)

func func1(done chan int) {
	fmt.Println("Hello world goroutine")
	for i := 0; i <= 100; i++ {
		done <- i
	}
}

func func2(done chan int) {
	for {
		v, _ := <-done
		if checkPrime(v) {
			fmt.Println(v, "is Prime")
		} else {
			fmt.Println(v, "is not Prime")
		}
	}
}

func checkPrime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	done := make(chan int)
	go func1(done)
	go func2(done)
	time.Sleep(5 * time.Second)
}
