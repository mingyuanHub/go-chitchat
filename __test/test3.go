package main

import "fmt"

func main() {
	// c := make(chan int)
	// c <- 1
	// <- c

	c := make(chan int, 3)
	c <- 2
	c <- 3
	c <- 4
	v1 := <- c
	fmt.Println(v1)
	v1 = <- c
	fmt.Println(v1)
	v1 = <- c
	fmt.Println(v1)
	c <- 5
}