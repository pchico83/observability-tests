package main

import (
	"net/http"
)

func createHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Only POST is allowed\n"))
		return
	}
	err := create()
	if err != nil {
		w.Write([]byte("KO\n"))
		return
	}
	w.Write([]byte("OK\n"))
}

func statusHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed\n"))
		return
	}
	err := status()
	if err != nil {
		w.Write([]byte("KO\n"))
		return
	}
	w.Write([]byte("OK\n"))
}

func restartHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Only POST is allowed\n"))
		return
	}
	err := restart()
	if err != nil {
		w.Write([]byte("KO\n"))
		return
	}
	w.Write([]byte("OK\n"))
}

func main() {
	http.HandleFunc("/create", createHTTP)
	http.HandleFunc("/status", statusHTTP)
	http.HandleFunc("/restart", restartHTTP)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
