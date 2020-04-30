package main

import (
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{
			Unid: cookie.Value
		}
		if ok, _ :=sess.Check(); !ok {
			err = error.New("Invalid session")
		}
	}
	return 
}

