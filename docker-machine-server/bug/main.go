package main

import (
	"net/http"
	"strings"
)

func createHTTP(w http.ResponseWriter, r *http.Request) {
	err := create()
	if err != nil {
		w.Write([]byte("KO"))
	}
	w.Write([]byte("OK"))
}

func statusHTTP(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func restartHTTP(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/create", createHTTP)
	http.HandleFunc("/status", statusHTTP)
	http.HandleFunc("/restart", restartHTTP)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
