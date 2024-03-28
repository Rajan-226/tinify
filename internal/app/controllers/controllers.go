package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/myProjects/tinify/internal/app/apis/tinify"
	"io"
	"net/http"
)

func Tinify(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	url, ok := data["url"].(string)
	if !ok {
		http.Error(w, "URL is missing or invalid", http.StatusBadRequest)
		return
	}

	tinifiedURL, err := tinify.Process(ctx, url, tinify.GetCore())
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to process request: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tinifiedURL))
	w.WriteHeader(http.StatusOK)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("redirect")
}

func Metrics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("metrics")
}
