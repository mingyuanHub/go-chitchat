package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/demo", demo)

	server.ListenAndServe()
}

func run() {
	fmt.Println("running...")
}

func demo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("demo")
}