package main

import (
	"fmt"
	"net/http"
)

const (
	FLAG = "CTF{congrats_you_used_HEAD}"
	HINT = "Hint: Sometimes the HEAD holds the key!"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodHead:
		// Add the flag to the headers for HEAD requests
		w.Header().Set("X-Flag", FLAG)
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		// Provide a hint for GET requests
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(HINT))
	default:
		// Respond with "Method Not Allowed" for other methods
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method Not Allowed"))
	}
}

func main() {
	http.HandleFunc("/", handler)

	port := 3000
	fmt.Printf("CTF challenge 'Get HEAD' running on http://localhost:%d/\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
