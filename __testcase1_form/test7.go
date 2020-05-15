package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<!DOCTYPE html>
<html>
<head>
	<title>Go Write</title>
</head>
<body>
hello world
</body>
</html>`
	
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "no response")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.baidu.com")
	w.WriteHeader(301)
}

type Post struct {
	User string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post {
		User: "mingyuan",
		Threads: []string{"first", "ming", "yuan"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/writeExample", writeExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}