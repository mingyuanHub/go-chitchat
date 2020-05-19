package main

import "fmt"
import "time"

func printlnNumbers2(w chan bool) {
	for i:= 0; i < 10; i ++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
	w <- true
}

func printLetters2(w chan bool) {
	for i:= 'A'; i < 'A' + 10; i ++ {
		time.Sleep(time.Microsecond)
		fmt.Printf("%c", i)
	}
	w <- true
}


func main() {
	// ch := make(chan int)
	// ch <- 1
	// i := <- ch
	// fmt.Println("i", i)

	w1, w2 := make(chan bool), make(chan bool)
	go printlnNumbers2(w1)
	go printLetters2(w2)

	a1 := <-w1
	a2 := <-w2

	fmt.Println(a1)
	fmt.Println(a2)
}