const express = require('express');
const path = require('path');
const fs = require('fs');
const app = express();

// Set EJS as view engine
app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'public'));  // EJS files should be in /public

// Serve static files (like CSS, JS, images)
app.use('/assets', express.static(path.join(__dirname, 'public/assets')));

// Home route
app.get('/', (req, res) => {
  res.render('index');
});

// LFI-vulnerable route
app.get('/page', (req, res) => {
  const page = req.query.page || 'home'; // default fallback
  const sanitizedPath = page.replace(/\.\./g, ''); // basic filter against directory traversal

  const fileToRender = path.join('pages', `${sanitizedPath}`); // e.g., 'pages/home'
  const fullPath = path.join(__dirname, 'public', `${fileToRender}.ejs`);

  if (fs.existsSync(fullPath)) {
    res.render(fileToRender); // render relative path (e.g., pages/home)
  } else {
    res.status(404).send('Page not found');
  }
});

// Start the server
const port = 3000;
app.listen(port, () => {
  console.log(`Server running on http://localhost:${port}`);
});
