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
  const query = `SELECT * FROM users WHERE username = ? AND password = ?`;

  db.get(query, [username, password], (err, row) => {
    if (err) {
      res.send('Error in SQL query.');
      return;
    }

    if (row) {
      res.send('<h2>Login successful! The flag is: secret_flag{12345}</h2><a href="/comment">Next Challenge</a>');
    } else {
      res.send('<h2>Invalid credentials!</h2>');
    }
  });
});

// XSS Vulnerability - Exploit to capture cookies or redirect
app.get('/comment', (req, res) => {
  res.send(`
    <h1>Leave a comment (XSS vulnerability)</h1>
    <form action="/comment" method="POST">
      Comment: <textarea name="comment"></textarea><br>
      <input type="submit" value="Submit">
    </form>
    <h2>Comments:</h2>
    <div id="comments">${req.query.comment ? req.query.comment : ''}</div>
  `);
});

app.post('/comment', (req, res) => {
  const { comment } = req.body;
  res.redirect('/comment?comment=' + comment);
});

// Insecure Direct Object Reference (IDOR) - Profile info leakage
app.get('/profile/:userId', (req, res) => {
  const userId = req.params.userId;
  const query = `SELECT * FROM users WHERE id = ?`;

  db.get(query, [userId], (err, row) => {
    if (err) {
      res.send('Error fetching user profile');
      return;
    }

    if (row) {
      res.send(`
        <h1>Profile of ${row.username}</h1>
        <p>ID: ${row.id}</p>
        <p>Username: ${row.username}</p>
        <p>Password: ${row.password}</p>
        <p>Flag for the next step: FLAG_2</p>
      `);
    } else {
      res.send('<h2>User not found</h2>');
    }
  });
});

// Command Injection - Execute arbitrary commands
app.get('/ping', (req, res) => {
  res.send(`
    <h1>Ping a Host</h1>
    <form action="/ping" method="POST">
      Host: <input type="text" name="host" /><br>
      <input type="submit" value="Ping" />
    </form>
  `);
});

app.post('/ping', (req, res) => {
  const { host } = req.body;
  exec(`ping ${host}`, (err, stdout, stderr) => {
    if (err || stderr) {
      res.send(`Error: ${stderr}`);
      return;
    }
    res.send(`<pre>${stdout}</pre>`);
  });
});

// File Upload Vulnerability - Upload reverse shell
const upload = multer({
  dest: 'uploads/',
  limits: { fileSize: 1 * 1024 * 1024 },
});

app.get('/upload', (req, res) => {
  res.send(`
    <h1>Upload a file (Vulnerable)</h1>
    <form action="/upload" method="POST" enctype="multipart/form-data">
      <input type="file" name="file" /><br>
      <input type="submit" value="Upload" />
    </form>
  `);
});

app.post('/upload', upload.single('file'), (req, res) => {
  if (!req.file) {
    return res.send('<h2>No file uploaded!</h2>');
  }
  res.send(`<h2>File uploaded: ${req.file.originalname}</h2>`);
});

// Weak Password Hashing (MD5) - Cracking the password hash
function hashPassword(password) {
  return crypto.createHash('md5').update(password).digest('hex');
}

app.get('/login-again', (req, res) => {
  res.send(`
    <h1>Login Again (Weak MD5 Hashing)</h1>
    <form action="/login-again" method="POST">
      Username: <input type="text" name="username"><br>
      Password: <input type="password" name="password"><br>
      <input type="submit" value="Login">
    </form>
  `);
});

app.post('/login-again', (req, res) => {
  const { username, password } = req.body;
  const hashedPassword = hashPassword(password);

  const query = `SELECT * FROM users WHERE username = ? AND password = ?`;

  db.get(query, [username, hashedPassword], (err, row) => {
    if (err) {
      res.send('Error in SQL query.');
      return;
    }

    if (row) {
      res.send('<h2>Login successful! The final flag is: FLAG_3</h2>');
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
