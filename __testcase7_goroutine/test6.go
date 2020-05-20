package main

import "fmt"

func callA(c chan string) {
	c <- "hello world"
	close(c)
}

func callB(c chan string) {
	c <- "hello ming"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callA(a)
	go callB(b)
	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		select {
		case msg , ok1 = <-a:
			if ok1 {
				fmt.Println("%s from A\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Println("%s from b\n", msg)
			}
		}
	}
}