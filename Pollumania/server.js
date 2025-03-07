const express = require('express');
const bodyParser = require('body-parser');
const session = require('express-session');
const fs = require('fs');
const path = require('path');

const app = express();
app.use(bodyParser.json());

// Session middleware to store user sessions
app.use(session({
    secret: 'ctf_secret_key', // Change this to a strong secret key
    resave: false,
    saveUninitialized: true,
    cookie: { secure: false } // Set to true if using HTTPS
}));

// Serve static files (like index.html) from the 'public' directory
app.use(express.static(path.join(__dirname, 'public')));

let users = {
    "guest": { username: "guest", role: "user" },
    "admin": { username: "admin", role: "admin" }
};

// Middleware for authentication
function isAuthenticated(req, res, next) {
    if (req.session.user && req.session.user.role === "admin") {
        return next();
    }
    return res.status(403).send("Access Denied");
}

// Login route
app.post('/login', (req, res) => {
    let { username } = req.body;
    if (users[username]) {
        req.session.user = users[username]; // Store user session
        res.json({ message: "Login successful", user: req.session.user });
    } else {
        res.status(401).json({ error: "Invalid credentials" });
    }
});

// Fixed route: Prevent prototype pollution
app.post('/update-profile', (req, res) => {
    let { username, role } = req.body;
    if (!username || !role) {
        return res.status(400).json({ error: "Invalid input" });
    }
    if (!users[username]) {
        return res.status(404).json({ error: "User not found" });
    }
    users[username].role = role;
    res.json({ message: "Profile updated!" });
});

// Admin-only route
app.get('/admin-panel', isAuthenticated, (req, res) => {
    try {
        let flag = fs.readFileSync("flag.txt", "utf8").trim();
        res.send(`Welcome Admin! Here is your flag: ${flag}`);
    } catch (err) {
        res.status(500).send("Flag file not found!");
    }
});

app.listen(3000, () => {
    console.log("Server running on port 3000");
});
