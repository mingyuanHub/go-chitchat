package main

import "fmt"
import "time"

func test() {
	fmt.Println(11111)
}

func test1() {
	fmt.Println(22222)
}

func main() {
	startTime := time.Now()
	defer func() {fmt.Println(time.Since(startTime))} ()

	time.Sleep(time.Second)
	defer test()

	fmt.Println("start")

	defer test1()

	for i := 0; i < 5; i ++ {
		defer fmt.Println(i)
	}

	
}