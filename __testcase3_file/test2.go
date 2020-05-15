package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	data := []byte("hello world")

	//第一种
	//写
	err := ioutil.WriteFile("data1",data, 0644)
	if err != nil {
		panic(err)
	}
	//读
	read1, _ := ioutil.ReadFile("data1")
	fmt.Println(string(read1))

	//第二种
	//写
	file1, _ := os.Create("data2")
	defer file1.Close()
	bytes, _ := file1.Write(data)
	fmt.Printf("wrote %d", bytes)
	//读
	file2, _ := os.Open("data2")
	defer file2.Close()
	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("read %d from file", bytes)
	fmt.Println(string(read2))
}
