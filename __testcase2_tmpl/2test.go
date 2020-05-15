package main

import (
 	"net/http"
 	"html/template"
 	// "math/rand"
 	// "time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t,_ := template.ParseFiles("2tmpl.html")
	// rand.Seed(time.Now().Unix())
	// t.Execute(w, rand.Intn(10) > 5)
	
	daysOfWeek := []string{"mon", "tue", "wed", "thu"}
	t.Execute(w, daysOfWeek)	
}

func process2(w http.ResponseWriter, r *http.Request) {
	t,_ := template.ParseFiles("2tmpl1.html", "2tmpl2.html")
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	server.ListenAndServe()
}