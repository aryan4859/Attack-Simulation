const express = require('express');
const bodyParser = require('body-parser');
const sqlite3 = require('sqlite3').verbose();
const dotenv = require('dotenv');
const multer = require('multer');
const crypto = require('crypto');
const { exec } = require('child_process');

dotenv.config();

const app = express();
app.use(bodyParser.urlencoded({ extended: true }));

// SQLite Database connection
const db = new sqlite3.Database('./database.db', (err) => {
  if (err) {
    console.error('Error opening database:', err.message);
  } else {
    console.log('Connected to the SQLite database.');

    // Create table if it doesn't exist
    db.run(`
      CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT,
        password TEXT
      );
    `, (err) => {
      if (err) {
        console.error('Error creating table:', err.message);
      } else {
        // Insert a sample user if the table is empty
        db.get('SELECT COUNT(*) AS count FROM users', (err, row) => {
          if (err) {
            console.error('Error checking users count:', err.message);
          } else if (row.count === 0) {
            // Insert sample user
            const query = 'INSERT INTO users (username, password) VALUES (?, ?)';
            db.run(query, ['admin', 'password123'], (err) => {
              if (err) {
                console.error('Error inserting sample user:', err.message);
              } else {
                console.log('Sample user inserted.');
              }
            });
          }
        });
      }
    });
  }
});

// Simulating a simple users table with a flag
// You must create this database and table manually before running the app if it doesn't already exist.

app.get('/', (req, res) => {
  res.send(`
    <h1>Welcome to the CTF Game!</h1>
    <form action="/login" method="POST">
      Username: <input type="text" name="username"><br>
      Password: <input type="password" name="password"><br>
      <input type="submit" value="Login">
    </form>
  `);
});

// SQL Injection - Login Bypass
app.post('/login', (req, res) => {
  const { username, password } = req.body;

  // ❌ Directly embedding user input in the query (VULNERABLE)
  const query = `SELECT * FROM users WHERE username = '${username}' AND password = '${password}'`;

  db.get(query, (err, row) => {
    if (err) {
      res.send('Error in SQL query.');
      return;
    }

    if (row) {
      res.send('<h2>Login successful! The flag is: secret_flag{12345}</h2>');
    } else {
      res.send('<h2>Invalid credentials!</h2>');
    }
  });
});


// Start server
const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`Server running on http://localhost:${PORT}`);
});
