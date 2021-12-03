package main

import (
	"fmt"
)

/*
	Task 8. Измените программу так, чтобы цифры от 1 до 9 печатались в консоль по порядку
*/
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go func() {
		for _, value := range []int{1, 4, 7} {
			<-c1
			fmt.Println(value)
			c2 <- 1
		}
	}()

	go func() {
		for _, value := range []int{2, 5, 8} {
			<-c2
			fmt.Println(value)
			c3 <- 1
		}
	}()

	go func() {
		for _, value := range []int{3, 6, 9} {
			<-c3
			fmt.Println(value)
			c1 <- 1
		}
	}()

	c1 <- 1

	var input string
	fmt.Scanln(&input)
}
