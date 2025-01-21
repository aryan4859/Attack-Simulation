package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

// The flag to simulate
const FLAG = "CTF{rc3_pwne6_200OK}"

// Serve the HTML frontend
func serveHTML(w http.ResponseWriter, r *http.Request) {
	// HTML content for the frontend
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>CTF RCE Challenge</title>
		</head>
		<body>
			<h1>Remote Code Execution (RCE) Challenge</h1>
			<form action="/cmd" method="POST">
				<label for="cmd">Enter Command (e.g., ping google.com):</label><br>
				<input type="text" id="cmd" name="cmd" required><br><br>
				<input type="submit" value="Execute">
			</form>
		</body>
		</html>
	`
	// Send the HTML response
	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write([]byte(html))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Only handle POST requests for command execution
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Use POST requests only."))
		return
	}

	// Parse user input (command)
	cmd := r.FormValue("cmd")
	if cmd == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Missing 'cmd' parameter."))
		return
	}

	// Block dangerous commands (like pwd)
	if strings.TrimSpace(cmd) == "pwd" {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("The 'pwd' command is blocked."))
		return
	}

	// Debugging: log command execution
	fmt.Printf("Executing command: %s\n", cmd)

	// Handle 'ping' command specifically
	if strings.HasPrefix(cmd, "ping") {
		host := strings.TrimPrefix(cmd, "ping ")
		if host == "" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing hostname for ping command."))
			return
		}

		// Execute the ping command
		out, err := exec.Command("ping", "-c", "4", host).Output()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Error: %s", err)))
			return
		}
		// Return ping output
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(out)
		return
	}

	// If the command isn't ping, treat it as a generic shell command
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	// Return the command output
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write(out)
}

func main() {
	// Handle the root endpoint to serve HTML
	http.HandleFunc("/", serveHTML)

	// Handle the /cmd endpoint for command execution
	http.HandleFunc("/cmd", handler)

	port := 3000
	fmt.Printf("RCE CTF challenge running on http://localhost:%d/\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
