package main

import (
	"net/http"
	"os/exec"
)

func createHTTP(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("ksync", "create")
	err := cmd.Run()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("OK"))
}

func getHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/create", createHTTP)
	http.HandleFunc("/get", getHTTP)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
