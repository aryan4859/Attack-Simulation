const express = require('express');
const bodyParser = require('body-parser');
const fs = require('fs');

const app = express();
app.use(bodyParser.json());

let users = {
    "guest": { username: "guest", role: "user" },
    "admin": { username: "admin", role: "admin" }
};

// Middleware for authentication
function isAuthenticated(req, res, next) {
    if (req.user && req.user.role === "admin") {
        return next();
    }
    return res.status(403).send("Access Denied");
}

// Simulated login
app.post('/login', (req, res) => {
    let { username } = req.body;
    if (users[username]) {
        req.user = users[username];
        res.json({ message: "Login successful", user: req.user });
    } else {
        res.status(401).json({ error: "Invalid credentials" });
    }
});

// Vulnerable route: Prototype Pollution
app.post('/update-profile', (req, res) => {
    Object.assign(users, req.body); // Merges request body into users object
    res.json({ message: "Profile updated!" });
});

// Admin-only route
app.get('/admin-panel', isAuthenticated, (req, res) => {
    let flag = fs.readFileSync("flag.txt", "utf8").trim();
    res.send(`Welcome Admin! Here is your flag: ${flag}`);
});

app.listen(3000, () => {
    console.log("Server running on port 3000");
});
