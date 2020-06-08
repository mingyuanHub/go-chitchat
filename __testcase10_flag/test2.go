package main

import (
	"flag"
	"fmt"
)

func main() {
	var age int
	flag.IntVar(&age, "age", 10, "set a age value")

	addr := flag.String("add", "shanghai", "set a localtion")
	flag.Parse()
	fmt.Println(age, *addr)
}
