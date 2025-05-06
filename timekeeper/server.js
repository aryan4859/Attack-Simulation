const express = require("express");
const path = require("path");
const app = express();

const PORT = 3000;
app.set("view engine", "ejs");
app.use(express.urlencoded({ extended: true }));
app.use(express.static("public"));

const users = [
  { username: "hannah", password: "tinkerbell", flag: "flag{T1m3K33p3r_U53r3nuM_3xp017}forge" }
];

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

app.get("/", (req, res) => {
  res.render("index", { error: null });
});

app.post("/login", async (req, res) => {
  const { username, password } = req.body;

  const user = users.find(u => u.username === username);

  // Introduce timing leak for user existence
  if (user) {
    await sleep(100); // Simulate checking password
    if (user.password === password) {
      return res.render("result", { msg: `Welcome, ${user.username}!`, flag: user.flag });
    } else {
      return res.render("index", { error: "Invalid credentials." });
    }
  } else {
    await sleep(5); // Shorter time if username doesn't exist
    return res.render("index", { error: "Invalid credentials." });
  }
});

app.listen(PORT, () => {
  console.log(`Timekeeper CTF running at http://localhost:${PORT}`);
});
