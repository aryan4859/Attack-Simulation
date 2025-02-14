package main

import (
	"fmt"
	"net/http"
	"os"
)

func iot1Handler(w http.ResponseWriter, r *http.Request) {
	// Open the .pcap file (e.g., abc.pcap) in the current directory
	file, err := os.Open("IOT1.bin")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file iot1
	w.Header().Set("Content-Disposition", "attachment; filename=IOT1.bin")
	w.Header().Set("Content-Type", "application/bin")

	// Serve the file
	http.ServeFile(w, r, "IOT1.bin")
}

func revHandler(w http.ResponseWriter, r *http.Request) { 
	file, err := os.Open("main")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file iot1
	w.Header().Set("Content-Disposition", "attachment; filename=main")
	w.Header().Set("Content-Type", "application/bin")

	// Serve the file
	http.ServeFile(w, r, "main")
}
func main() {
	// Set up the iot1 handler
	http.HandleFunc("/iot1", iot1Handler) 
	http.HandleFunc("/main", revHandler)

	// Start the server on port 8080
	fmt.Println("Server is running at http://localhost:8080/")
	err := http.ListenAndServe(":3003", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
