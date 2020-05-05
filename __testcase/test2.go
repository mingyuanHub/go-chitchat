package main

import (
 	"fmt"
 	"net/http"   
)

func hand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s!", 123)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", hand)

	serve := http.Server{
		Addr: "127.0.0.1:8080",
		Handler:mux,
	}

	serve.ListenAndServe()
}