package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Get the PORT from environment variables, default to 3030 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/read", readFileHandler)

	fmt.Printf("Server started on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<h1>Welcome to PathFinder!</h1>
		<p>A poorly designed Go web application allows users to view files based on a URL parameter. Your task is to exploit the vulnerability to include and read sensitive files from the server, ultimately retrieving the flag.</p>
		<style>
        ul {
            list-style-type: none;
            font-family: Arial, sans-serif;
        }
        .folder::before {
            content: "📂 ";
        }
        .file::before {
            content: "📄 ";
        }
    </style>   
    <ul>
        <li class="folder">path-transversal/
            <ul>
                <li class="file">main.go</li>
                <li class="folder">files/
                    <ul>
                        <li class="file">about.txt</li>
                    </ul>
                </li>
                <li class="file">flag.txt</li>
            </ul>
        </li>
    </ul> 
		<a href="read?file=about.txt">Learn more about</a>
	`))
}

func readFileHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")

	if file == "" {
		http.Error(w, "No file specified", http.StatusBadRequest)
		return
	}

	// Simulate an insecure file inclusion vulnerability
	path := filepath.Join("./files", file)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
