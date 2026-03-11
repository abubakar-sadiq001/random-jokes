package main

import (
	"fmt"
	"io"
	"net/http"
)

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/api/jokes", corsMiddleware(http.HandlerFunc(apiHandler)))

	fmt.Println("🚀 Server running on :8080")
	http.ListenAndServe("localhost:8080", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		http.Error(w, "Failed to fetch joke", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Check if the external API returned an error status
	if response.StatusCode != http.StatusOK {
		http.Error(w, "Joke API returned error", response.StatusCode)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Failed to read joke", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(body)
}
