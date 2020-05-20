package main

import "fmt"
import "time"

func ex(w chan int) {
	w <- 2
}

func ax(w chan int) {
	n := <-w
	fmt.Println(n)
}

func main() {
	c := make(chan int)
	go ex(c)
	go ax(c)
	time.Sleep(100 * time.Millisecond)

}