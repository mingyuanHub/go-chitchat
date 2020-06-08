package main

import "os"
import "fmt"

func main() {
	if len(os.Args) > 0 {
		for index, args := range os.Args {
			fmt.Println(index, args)
		}
	}
}