package main

import "fmt"

func main(){
	i := new(int)
	a := 3
	i = &a
	fmt.Println(*i)
}