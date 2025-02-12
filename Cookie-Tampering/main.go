package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// Index handler that checks for the Base64-encoded user cookie
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Try to get the "user" cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		// If no cookie exists, set default "guest" cookie
		encodedGuest := base64.StdEncoding.EncodeToString([]byte("guest"))
		http.SetCookie(w, &http.Cookie{
			Name:  "user",
			Value: encodedGuest,
			Path:  "/",
		})
		fmt.Fprintln(w, "No cookie found. Defaulting to guest. Try modifying your cookie.")
		return
	}

	// Decode the Base64 value
	decodedValue, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid cookie encoding!", http.StatusBadRequest)
		return
	}

	// Check if the decoded value is "admin"
	if string(decodedValue) == "admin" {
		fmt.Fprintln(w, "Welcome, Admin! 🎉 Here is your flag: Flag{C00K13_7839ER15g_365c0229669ea435}")
	} else {
		fmt.Fprintf(w, "Welcome, %s! Try modifying your cookie to become an admin.", decodedValue)
	}
}

// Login handler that sets the "user" cookie with Base64-encoded "guest"
func loginHandler(w http.ResponseWriter, r *http.Request) {
	encodedGuest := base64.StdEncoding.EncodeToString([]byte("guest"))
	http.SetCookie(w, &http.Cookie{
		Name:  "user",
		Value: encodedGuest,
		Path:  "/",
	})
	fmt.Fprintln(w, "Logged in as Guest. Try modifying your cookie!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
