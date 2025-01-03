package main

import (
	"fmt"
	"net/http"
	"strconv" 
)

var profiles = map[int]string{
	1: "Alice: Basic User Profile",
2: "Bob: Basic User Profile",
3: "Charlie: Basic User Profile",
4: "David: Basic User Profile",
5: "Eve: Basic User Profile",
6: "Frank: Basic User Profile",
7: "Grace: Basic User Profile",
8: "Hannah: Basic User Profile",
9: "Ivy: Basic User Profile",
10: "Jack: Basic User Profile",
11: "Kathy: Basic User Profile",
12: "Liam: Basic User Profile",
13: "Mona: Basic User Profile",
14: "Nathan: Basic User Profile",
15: "Olivia: Basic User Profile",
16: "Paul: Basic User Profile",
17: "Quincy: Basic User Profile",
18: "Rita: Basic User Profile",
19: "Sam: Basic User Profile",
20: "Tina: Basic User Profile",
21: "Ursula: Basic User Profile",
22: "Victor: Basic User Profile",
23: "Wendy: Basic User Profile",
24: "Xander: Basic User Profile",
25: "Yara: Basic User Profile",
26: "Zane: Basic User Profile",
27: "Abigail: Basic User Profile",
28: "Ben: Basic User Profile",
29: "Catherine: Basic User Profile",
30: "Daniel: Basic User Profile",
31: "Eva: Basic User Profile",
32: "Felix: Basic User Profile",
33: "Gina: Basic User Profile",
34: "Harry: Basic User Profile",
35: "Isla: Basic User Profile",
36: "James: Basic User Profile",
37: "Kimberly: Basic User Profile",
38: "Luca: Basic User Profile",
39: "Maya: Basic User Profile",
40: "Noah: Basic User Profile",
41: "Ophelia: Basic User Profile",
42: "Peter: Basic User Profile",
43: "Quinn: Basic User Profile",
44: "Riley: VIP User Profile",
45: "Sophia: Basic User Profile",
46: "Travis: Basic User Profile",
47: "Ulysses: Basic User Profile",
48: "Vera: Basic User Profile",
49: "Will: Basic User Profile",
50: "Xena: Basic User Profile",
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the 'user_id' from the query parameters
	userIDParam := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDParam)

	// If the user_id is missing or invalid, return an error
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid or missing user_id parameter", http.StatusBadRequest)
		return
	}

	// Simulate fetching the profile from a "database"
	profile, exists := profiles[userID]
	if !exists {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	// Display the profile
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>IDOR Challenge</title>
		</head>
		<body>
			<h1>IDOR Challenge</h1>
			<p>Welcome to your profile page!</p>
			<h2>Your Profile:</h2>
			<p>%s</p>
			<hr>
			<div style="display:none;" id="flag">
				Congratulations! Your flag is: FLAG{IDOR_VULNERABILITY_EXPLOITATION}
			</div>
		</body>
		</html>
	`, profile)

	// If the user accesses a specific profile (like User 3), display the flag
	if userID == 44 {
		fmt.Fprintf(w, `<script>document.getElementById('flag').style.display = 'block';</script>`)
	}
}

func main() {
	http.HandleFunc("/profile", profileHandler)

	// Start the server on port 8080
	fmt.Println("Server is running at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
