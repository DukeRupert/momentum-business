package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Get configuration from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Support multiple origins (comma-separated)
	allowedOriginsStr := os.Getenv("ALLOWED_ORIGINS")
	if allowedOriginsStr == "" {
		allowedOriginsStr = "http://localhost:1313"
	}

	// Parse origins into a map for fast lookup
	allowedOrigins := make(map[string]bool)
	for _, origin := range strings.Split(allowedOriginsStr, ",") {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			allowedOrigins[origin] = true
		}
	}

	// Create router
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("POST /api/contact", handleContact)
	mux.HandleFunc("GET /api/health", handleHealth)

	// Wrap with CORS middleware
	handler := corsMiddleware(mux, allowedOrigins)

	log.Printf("Starting API server on port %s", port)
	log.Printf("Allowed origins: %v", allowedOriginsStr)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}

func corsMiddleware(next http.Handler, allowedOrigins map[string]bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Check if the origin is allowed
		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
