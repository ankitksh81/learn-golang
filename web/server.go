package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

const (
	HOST = "localhost"
	PORT = "8080"
	USER = "admin"
	PASS = "admin"
)

/*---------------------------
	HTTP SERVER WITH AUTH
----------------------------*/

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
	// w.Write([]byte("Welcome!")) // another way to print using response writer
}

// Basic Authentication
func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user),
			[]byte(USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass),
			[]byte(PASS)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the site.\n"))
			return
		}
		handler(w, r)
	}
}

func main() {
	http.HandleFunc("/", BasicAuth(index, "Please enter your username & password"))
	// http.HandleFunc("/", index)

	// server
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
