package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func main() {
	setupLogger()

	// Get configuration from environment
	// Use API_PORT to avoid conflict with Caddy's PORT
	port := os.Getenv("API_PORT")
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

	// Wrap with CORS and access logging middleware
	handler := loggingMiddleware(corsMiddleware(mux, allowedOrigins))

	slog.Info("server starting", "addr", ":"+port, "allowed_origins", allowedOriginsStr)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		slog.Error("server stopped", "addr", ":"+port, "error", err.Error())
		os.Exit(1)
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
