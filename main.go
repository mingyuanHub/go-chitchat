package main

import (
	"net/http"
	"html/template"
	"fmt"
	"go-chitchat/data"
	"go-chitchat/conf"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	files := []string{
// 		"templates/layout.html",
// 		"templates/navbar.html",
// 		"templates/index.html",
// 	}

// 	templates := template.Must(template.ParseFiles(files...))
// 	threads, err := data.Threads()
// 	if err == nil {
// 		templates.ExecuteTemplate(w, "layout", threads)
// 	}
// }

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _ , file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html"), file)
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := conf.util.session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("D:/Go/www/go-chitchat/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server {
		Addr: "0.0.0.0:8000",
		Handler: mux,
	}

	server.ListenAndServe()
}