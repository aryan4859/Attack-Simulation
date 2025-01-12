package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

const FLAG = "CTF{rc3_pwne6_200OK}"

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Use POST requests only."))
		return
	}

	// Parse user input
	cmd := r.URL.Query().Get("cmd")
	if cmd == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Missing 'cmd' parameter."))
		return
	}

	// Debugging: print out the command being executed
	fmt.Printf("Executing command: %s\n", cmd)

	// Test with simple commands first
	if strings.TrimSpace(cmd) == "echo test" {
		_, _ = w.Write([]byte("test"))
		return
	}

	// Simulate insecure command execution
	// WARNING: This is intentionally vulnerable; avoid using this in real applications.
	out, err := exec.Command("sh", "-c", cmd).Output()

	if err != nil {
		// Debugging: print out the error message
		fmt.Printf("Error executing command: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	// Simulating the flag file for testing purposes
	if strings.TrimSpace(cmd) == "cat flag" {
		// Read the local flag file
		_, _ = w.Write([]byte(FLAG))
		return
	}

	// Return the command output
	_, _ = w.Write(out)
}

func main() {
	http.HandleFunc("/", handler)

	port := 3000
	fmt.Printf("RCE CTF challenge running on http://localhost:%d/\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
