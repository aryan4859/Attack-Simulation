package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

const css = `
<style>
    body { font-family: Arial, sans-serif; text-align: center; padding: 50px; background-color: #f4f4f4; }
    h2 { color: #333; }
    p { color: #666; }
    a { display: inline-block; padding: 10px 20px; color: white; background-color: #007BFF; text-decoration: none; border-radius: 5px; }
    a:hover { background-color: #0056b3; }
</style>
`

// indexHandler serves the home page with user authentication status
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
		html := `
		<!DOCTYPE html>
		<html>
		<head><title>Welcome</title>` + css + `</head>
		<body>
			<h2>No cookie found. Defaulting to guest.</h2>
			<p>Try modifying your cookie to see what happens!</p>
			<a href="/login">Login</a>
		</body>
		</html>`
		fmt.Fprintln(w, html)
		return
	}

	// Decode the Base64 value
	decodedValue, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid cookie encoding!", http.StatusBadRequest)
		return
	}

	// Render HTML based on user role
	var html string
	if string(decodedValue) == "admin" {
		html = `
		<!DOCTYPE html>
		<html>
		<head><title>Admin Panel</title>` + css + `</head>
		<body>
			<h2>Welcome, Admin! 🎉</h2>
			<p>Here is your flag: <strong>flag{C00K13_7839ER15g_365c0229669ea435}</strong></p>
		</body>
		</html>`
	} else {
		html = fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head><title>Welcome</title>`+css+`</head>
		<body>
			<h2>Welcome, %s!</h2>
			<p>Try modifying your cookie to become an admin.</p>
			<a href="/login">Re-login as Guest</a>
		</body>
		</html>`, decodedValue)
	}
	fmt.Fprintln(w, html)
}

// loginHandler sets the "user" cookie to "guest"
func loginHandler(w http.ResponseWriter, r *http.Request) {
	encodedGuest := base64.StdEncoding.EncodeToString([]byte("guest"))
	http.SetCookie(w, &http.Cookie{
		Name:  "user",
		Value: encodedGuest,
		Path:  "/",
	})

	html := `
	<!DOCTYPE html>
	<html>
	<head><title>Login</title>` + css + `</head>
	<body>
		<h2>Logged in as Guest</h2>
		<p>Try modifying your cookie!</p>
		<a href="/">Go to Home</a>
	</body>
	</html>`

	fmt.Fprintln(w, html)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
