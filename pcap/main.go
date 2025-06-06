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
	file, err := os.Open("garden.jpg")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=garden.jpg")
	w.Header().Set("Content-Type", "image/jpg")

	// Serve the file
	http.ServeFile(w, r, "garden.jpg")
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

func wavHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("secret.wav")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=secret.wav")
	w.Header().Set("Content-Type", "audio/wav")

	// Serve the file
	http.ServeFile(w, r, "secret.wav")
}

func secretsHandler(w http.ResponseWriter, r *http.Request) {
	// Open the secrets file
	file, err := os.Open("cipher.txt")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=cipher.txt")
	w.Header().Set("Content-Type", "text/plain")

	// Serve the file
	http.ServeFile(w, r, "cipher.txt")
}

func alfridaHandler(w http.ResponseWriter, r *http.Request){
	file, err := os.Open("alfrida.apk")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=alfrida.apk")
	w.Header().Set("Content-Type", "text/plain")

	// Serve the file
	http.ServeFile(w, r, "alfrida.apk")
}

func main() {
	// Set up the download handler
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/garden", gardenHandler)
	http.HandleFunc("/flag", flagHandler)
	http.HandleFunc("/state", stateHandler)
	http.HandleFunc("/fast", fastHandler)
	http.HandleFunc("/xoracle", xoracleHandler)
	http.HandleFunc("/wav", wavHandler)
	http.HandleFunc("/secrets", secretsHandler)
	http.HandleFunc("/alfrida", alfridaHandler)

	// Start the server on port 8080
	fmt.Println("Server is running at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
