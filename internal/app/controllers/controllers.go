package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Rajan-226/tinify/internal/app/tinify"
	"github.com/Rajan-226/tinify/internal/pkg/utils"
	"github.com/gorilla/mux"
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

	if !utils.IsValidURL(url) {
		http.Error(w, "URL is invalid", http.StatusBadRequest)
		return
	}

	tinifiedURL, err := tinify.Create(ctx, url, tinify.GetCore())
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to process request: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tinifiedURL))
	w.WriteHeader(http.StatusOK)
}

func GetAllURLs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	longToShortURL, err := tinify.GetAllUrls(ctx, tinify.GetCore())
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to process request: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(longToShortURL)
	if err != nil {
		http.Error(w, "Failed to marshal longToShortURL map", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	shortURI := vars["path"]
	if shortURI == "" {
		http.Error(w, "Invalid short url", http.StatusBadRequest)
		return
	}

	longURL, err := tinify.Redirect(ctx, shortURI, tinify.GetCore())
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to process request: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func Metrics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	topDomains, err := tinify.GetCore().GetTopShortenedDomains(ctx, 3)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve metrics: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(topDomains)
	if err != nil {
		http.Error(w, "Failed to encode metrics", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("Failed to write response:", err)
	}
}
