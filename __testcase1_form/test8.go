package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie {
		Name:  "firstCookie",
		Value: "go web cookie1",
		HttpOnly: true,
	}

	c2 := http.Cookie {
		Name:  "secondCookie",
		Value: "go web cookie2",
		HttpOnly: true,
	}

	// w.Header().Set("Set-Cookie", c1.String())
	// w.Header().Add("Set-Cookie", c2.String())

	http.SetCookie(w, &c1)  //传入指针
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)

	c1, err := r.Cookie("firstCookie")

	if err != nil {
	}

	cs := r.Cookies()

	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main() {
	server  := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/cookie", setCookie)
	http.HandleFunc("/getcookie", getCookie)

	server.ListenAndServe()
}