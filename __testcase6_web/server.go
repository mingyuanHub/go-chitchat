package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	// "goChitchat/__testcase6_web"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	// fmt.Println(body)
	r.Body.Read(body)
	// fmt.Println(body)
	var post Post
	json.Unmarshal(body, &post)
	err = post.Create()
	fmt.Println(post.Id)
	if err != nil {
		return
	}
	response200(w)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	post, err := retrieve(id)
	if err != nil {
		return
	}
	json.Unmarshal(body, &post)
	err = post.Update()
	if err != nil {
		return
	}

	response200(w)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.Delete()
	if err != nil {
		return
	}
	response200(w)
	return
}

func response200(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintln(w, "{\"code\":200}")
}
