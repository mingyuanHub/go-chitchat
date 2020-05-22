package main

import "fmt"

func test(c chan int) {
	c <- 1
	c <- 2
}

func test2(c chan int) {
	a := <- c
	b := <- c
	fmt.Println(a, b)
}

func main() {
	c := make(chan int)
	go test(c)
	go test2(c)
}