package main

import (
	"net/http"
	"html/template"
	"math/rand"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("5tmpl1.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func color(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	var t *template.Template

	if rand.Intn(10) > 7 {
		t, _ = template.ParseFiles("5tmpl1.html", "5tmpl_red.html")
	} else if rand.Intn(10) > 4 {
		t, _ = template.ParseFiles("5tmpl1.html", "5tmpl_blue.html")
	} else {
		t, _ = template.ParseFiles("5tmpl1.html")
	}

	t.ExecuteTemplate(w, "layout", nil)

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/color", color)
	server.ListenAndServe()
}