package main

import (
	"fmt"
	"net/http"
	"os"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// Open the .pcap file (e.g., abc.pcap) in the current directory
	file, err := os.Open("abc.pcap")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=abc.pcap")
	w.Header().Set("Content-Type", "application/vnd.tcpdump.pcap")

	// Serve the file
	http.ServeFile(w, r, "abc.pcap")
}

func gardenHandler(w http.ResponseWriter, r *http.Request) { 
	file, err := os.Open("garden(1).jpg")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=garden(1).jpg")
	w.Header().Set("Content-Type", "image/jpg")

	// Serve the file
	http.ServeFile(w, r, "garden(1).jpg")
}

func main() {
	// Set up the download handler
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/garden", gardenHandler)

	// Start the server on port 8080
	fmt.Println("Server is running at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
