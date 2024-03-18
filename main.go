package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	centralmux := http.NewServeMux()
	centralmux.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", centralmux)
	if err != nil {
		fmt.Println("Error", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.Path

	switch endpoint {
	case "/login":
		handlerLogin(w, r)
	case "/signup":
		handlerSignup(w, r)
	default:
		http.NotFound(w, r)
	}
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Info: User Login Info")
}

func handlerSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Info: User Signup Info")
}
