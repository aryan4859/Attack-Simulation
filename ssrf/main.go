package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const FLAG = "FLAG{SSRF_3xP0s3d_3000}" // Hidden flag

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Sneaky Requests! Try to access the admin panel.")
	})

	http.HandleFunc("/fetch", fetchHandler)
	http.HandleFunc("/admin", adminHandler)

	fmt.Println("Server started on :3000")
	http.ListenAndServe(":3000", nil)
}

func fetchHandler(w http.ResponseWriter, r *http.Request) {
	queryURL := r.URL.Query().Get("url")
	if queryURL == "" {
		http.Error(w, `{"error": "URL parameter is required"}`, http.StatusBadRequest)
		return
	}

	// Parse and validate URL
	parsedURL, err := url.Parse(queryURL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		http.Error(w, `{"error": "Invalid URL"}`, http.StatusBadRequest)
		return
	}

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(queryURL) // Server fetches the URL
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch URL"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read only the first 300 bytes
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 300))

	responseData := map[string]interface{}{
		"status":  resp.StatusCode,
		"content": string(body),
	}

	jsonResponse, _ := json.Marshal(responseData)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow requests from localhost
	if r.Host != "localhost:3000" && r.Host != "127.0.0.1:3000" {
		http.Error(w, "403 Forbidden: Access denied", http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "Admin Panel: The flag is %s", FLAG)
}
