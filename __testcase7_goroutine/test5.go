package main

import "fmt"
import "time"

func callerA(c chan string) {
	c <- "hello world"
}

func callerB(c chan string) {
	c <- "hello mmy"
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond)
		select {
		case msg:= <- a:
			fmt.Println(msg)
		case msg:= <- b:
			fmt.Println(msg)
		default:
			fmt.Println("13")
		}
	}
}