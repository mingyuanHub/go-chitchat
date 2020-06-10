package main

import "fmt"

func main() {
	defer func(){
		fmt.Println(4444)
		recover()
		fmt.Println(555)
	}()

	fmt.Println(2222)

	panic("panic 123123")

	fmt.Println(33333)
}