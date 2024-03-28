package controllers

import (
	"fmt"
	"net/http"
)

func Tinify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("tinify")

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("OK"))
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("redirect")
}

func Metrics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("metrics")
}
