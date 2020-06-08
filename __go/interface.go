package main

import "fmt"
import "reflect"

type Duck interface {
	Walk()
	Quack()
}

type Cat struct {

}

func (c *Cat) Walk() {
	fmt.Println(" * walk")
}

func (c *Cat) Quack() {
	fmt.Println(" * quack")
}

// func (c Cat) Walk() {
// 	fmt.Println("s walk")
// }

// func (c Cat) Quack() {
// 	fmt.Println("s quack")
// }

func main() {
	c := Cat{}
	c.Walk()
	c.Quack()

	d := reflect.ValueOf(1)
	fmt.Println(d)

	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))
	fmt.Println("ValueOf author:", reflect.TypeOf(reflect.ValueOf(author).Interface()))
	

	hash := map[string]int {
		"1":1,
		"2":2,
		"3":3,
	}

	for i, v := range(hash) {
		fmt.Println(i + ":", v)
	}
}