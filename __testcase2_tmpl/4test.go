package main

import (
	"net/http"
	"html/template"
)

//自行转义
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("4tmpl.html")
	t.Execute(w, r.FormValue("comment"))
}

//不做转义
func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("4tmpl.html")
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("4tmpl2.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	http.HandleFunc("/form", form)

	server.ListenAndServe()
}
