package main

import (
	"fmt"
)

func init() {
	fmt.Println(111)
}

func main() {
	var a = 1;

	fmt.Println(&a)

	fmt.Println(*&a)
}