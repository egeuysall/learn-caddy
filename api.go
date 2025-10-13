package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// This is the JSON structure we'll send back
type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Create our response
		response := Response{
			Message: "Hello from Go API!",
			Status:  "running",
		}

		// Send it as JSON
		json.NewEncoder(w).Encode(response)
	})

	// Start the server on port 8080
	log.Println("🚀 Go API running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
