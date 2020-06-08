package main

import (
	"flag"
	"fmt"
)

func main() {
	var ip = flag.Int("flag", 1, "get a flag value")

	var ip1 int
	flag.IntVar(&ip1, "ip1", 12, "set ip1 value")

	flag.Parse()

	fmt.Println(*ip)
	fmt.Println(ip1)

	flag.PrintDefaults()
}
