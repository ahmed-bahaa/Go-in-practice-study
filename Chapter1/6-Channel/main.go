package main

import (
	"fmt"
	"time"
)

func printCoun(c chan int) {
	num := 0

	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)
	go printCoun(c)
	myList := []int{1, 2, 3, -1}
	for _, v := range myList {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of Program")
}
