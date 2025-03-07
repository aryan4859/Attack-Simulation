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

func flagHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("flag.png")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=flag.png")
	w.Header().Set("Content-Type", "image/png")

	// Serve the file
	http.ServeFile(w, r, "flag.png")
}

func stateHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("state.jpg")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=state.jpg")
	w.Header().Set("Content-Type", "image/jpg")

	// Serve the file
	http.ServeFile(w, r, "state.jpg")
}

func fastHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("file3.png")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=file3.png")
	w.Header().Set("Content-Type", "image/png")

	// Serve the file
	http.ServeFile(w, r, "file3.png")
}

func xoracleHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("main")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=main")
	w.Header().Set("Content-Type", "executable")

	// Serve the file
	http.ServeFile(w, r, "main")
}

func main() {
	// Set up the download handler
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/garden", gardenHandler)
	http.HandleFunc("/flag", flagHandler)
	http.HandleFunc("/state", stateHandler)
	http.HandleFunc("/fast", fastHandler)
	http.HandleFunc("/xoracle", xoracleHandler)

	// Start the server on port 8080
	fmt.Println("Server is running at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
