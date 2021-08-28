/*---------------------------
	HTTP SERVER WITH AUTH
----------------------------*/

package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
)

// Reading from .env file
func viperEnv(key string) string {
	viper.SetConfigFile(".env") // to set the config file
	err := viper.ReadInConfig() // to read from the config set above
	if err != nil {
		log.Fatalf("Error reading from config file: %v", err)
	}

	value, ok := viper.Get(key).(string) // type assertion value from interface -> string
	if !ok {
		log.Fatalf("Invalid type")
	}

	return value
}

// Environment variables
var HOST = viperEnv("HOST")
var PORT = viperEnv("PORT")
var USER = viperEnv("USER")
var PASS = viperEnv("PASS")

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

// func main() {
// 	http.HandleFunc("/", BasicAuth(index, "Please enter your username & password"))
// 	// http.HandleFunc("/", index)

// 	// server
// 	err := http.ListenAndServe(HOST+":"+PORT, nil)
// 	if err != nil {
// 		log.Fatalf("error starting server: %v", err)
// 	}
// }

/*----------------------------------------------
	HTTP SERVER WITH MUX WITH GZIP COMPRESSION
----------------------------------------------*/

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	err := http.ListenAndServe(HOST+":"+PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
