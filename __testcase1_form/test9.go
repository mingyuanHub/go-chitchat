package main 

import (
  	"fmt"
  	"net/http"
  	"time"
  	"encoding/base64"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	msg := []byte("hello")
	c := http.Cookie{
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}

	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")

	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "no message found")
		}
	} else {
		rc := http.Cookie{
			Name: "flash",
			MaxAge: -1,
			Expires: time.Unix(1, 0),
		}

		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/get", showMessage)

	server.ListenAndServe()
}