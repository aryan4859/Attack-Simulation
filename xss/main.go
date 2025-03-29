package main

import (
	"fmt"
	"net/http"
	"time"
)

// Handler for the search page
func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Set a session cookie (simulating user login)
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   "THM{refl3ct3d_x55_victim_fla9}",
		Path:    "/",
		Expires: time.Now().Add(1 * time.Hour),
	})

	// Serve the search page with the reflected input (XSS vulnerability)
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Search - CrossXReflection</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            padding: 20px;
            background-color: #f4f4f9;
        }
        h1 {
            color: #333;
            text-align: center;
        }
        form {
            margin-top: 20px;
            text-align: center;
        }
        input[type="text"] {
            padding: 10px;
            width: 250px;
            font-size: 16px;
            margin-right: 10px;
        }
        button {
            padding: 10px 20px;
            background-color: #3498db;
            color: white;
            border: none;
            font-size: 16px;
            cursor: pointer;
            border-radius: 5px;
        }
        button:hover {
            background-color: #2980b9;
        }
        .greeting {
            margin-top: 20px;
            color: #27ae60;
            font-size: 18px;
        }
        .container {
            max-width: 900px;
            margin: 0 auto;
            padding: 20px;
            background-color: white;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            border-radius: 8px;
        }
        .flag-info {
            margin-top: 30px;
            font-size: 14px;
            color: #e74c3c;
        }
        .hidden-flag {
            display: none;
        }
        .hidden-info {
            font-size: 14px;
            color: #34495e;
        }
        .info-header {
            font-size: 18px;
            margin-top: 30px;
        }
    </style>
</head>
<body>

    <div class="container">
        <h1>Welcome to the Xreflection</h1>

        <form action="/" method="GET">
            <label for="query">Enter your search term:</label><br><br>
            <input type="text" id="query" name="query" required>
            <button type="submit">Search</button>
        </form>

        <div class="greeting">
            <p>Search results for: `)

	// Reflect the "query" query parameter directly back to the page (vulnerable to XSS)
	query := r.URL.Query().Get("query")
	if query != "" {
		// Here we reflect back the search input without any validation, enabling XSS
		fmt.Fprintf(w, "%s", query)
	}

	fmt.Fprintf(w, `</p>

        <div class="hidden-info">
            <p>This is a custom hidden element for added complexity. It contains a lot of random text that could be a part of a larger application flow.</p>
        </div>

        
        <div class="info-header">
            <p>Some useful information might be hidden below.</p>
        </div>

    </div>

    <script>
        // Just to make it harder for users, add some distraction
        document.querySelector('.hidden-info').innerHTML += "<p>Some random distraction text that makes it harder to focus on the flag.</p>";

        // Add some DOM manipulation to increase complexity
        setTimeout(function() {
            document.querySelector('.hidden-flag').style.display = "block";
        }, 3000);
    </script>

</body>
</html>`)
}

// Main function to start the Go server
func main() {
	http.HandleFunc("/", searchHandler)
	http.ListenAndServe(":4900", nil)
}
