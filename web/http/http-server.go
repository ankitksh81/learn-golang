/*---------------------------
	HTTP SERVER WITH AUTH
----------------------------*/

package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

/*
func main() {
	http.HandleFunc("/", BasicAuth(index, "Please enter your username & password"))
	// http.HandleFunc("/", index)

	// server
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
*/

/*--------------------------------------------
	HTTP SERVER WITH MUX WITH GZIP COMPRESSION
----------------------------------------------*/
/*
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	err := http.ListenAndServe(HOST+":"+PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
*/

/*----------------------
	HTTP REQUEST ROUTING
-----------------------*/

/*
// index
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
	// w.Write([]byte("Welcome!")) // another way to print using response writer
}
*/

// login handler
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Page!")
}

// logout handler
func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Page!")
}

/*
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
		return
	}
}
*/

/*----------------------------------------
	HTTP REQUEST ROUTING USING GORILLA MUX
------------------------------------------*/
var GetRequestHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
})

var PostRequestHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's a Post request!"))
})

// dynamic requests handler
var PathVariableHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Write([]byte("Hi " + name + "\n"))
})

func main() {
	router := mux.NewRouter()
	router.Handle("/", GetRequestHandler).Methods("GET")
	router.Handle("/post", PostRequestHandler).Methods("POST")
	router.Handle("/hello/{name}", PathVariableHandler).Methods("GET", "PUT")

	http.ListenAndServe(HOST+":"+PORT, router)
}
