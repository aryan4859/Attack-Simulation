// server.js
// BizLogic Bypass CTF challenge with embedded single-file HTML frontend
// Run with: node server.js
// The server serves both the API and a simple frontend at /

const express = require("express");
const bodyParser = require("body-parser");
const jwt = require("jsonwebtoken");
const fs = require("fs");
const path = require("path");

const app = express();
app.use(bodyParser.json());

/*
  Simple in-memory "database"
  - users: {id, username, role, managerId}
  - expenses: {id, submitterId, amount, status, approvalToken}
*/
let users = [
  { id: 1, username: "alice", role: "employee", managerId: 2 },
  { id: 2, username: "bob", role: "manager", managerId: null },
  { id: 3, username: "charlie", role: "finance", managerId: null },
];

let expenses = [];
let nextExpenseId = 1;

// Intentionally weak secret and predictable token generation logic
const JWT_SECRET = "acmecorp-secret"; // <-- weak / known in challenge
const PORT = 3000;

// Helper: simple auth middleware that trusts X-User header (intentionally)
function auth(req, res, next) {
  // In this challenge the developer "optimized" local dev convenience
  // by allowing client-side header to set the user id. This is the core of the bug.
  // In the real world: never trust client-supplied user identity.
  const uid = parseInt(req.header("X-User") || "0", 10);
  const user = users.find((u) => u.id === uid);
  if (!user) return res.status(401).json({ error: "Invalid user" });
  req.user = user;
  next();
}

// Login (convenience route): returns a JWT but not used by server for auth
app.post("/login", (req, res) => {
  const { username } = req.body;
  const user = users.find((u) => u.username === username);
  if (!user) return res.status(401).json({ error: "No such user" });

  const token = jwt.sign(
    { id: user.id, username: user.username, role: user.role },
    JWT_SECRET,
    { expiresIn: "1h" }
  );

  res.json({
    token,
    message: "Use X-User header for the challenge (dev override).",
  });
});

// Submit expense
// Submit expense
app.post("/expenses", auth, (req, res) => {
  const { amount, description } = req.body;
  if (!amount) return res.status(400).json({ error: "amount required" });

  const id = nextExpenseId++;
  const approvalToken = jwt.sign(
    { expenseId: id, managerId: req.user.managerId },
    JWT_SECRET
  );

  const expense = {
    id,
    submitterId: req.user.id,
    amount,
    description,
    status: "submitted",
    approvalToken,
  };
  expenses.push(expense);

  // Handle users with no manager
  let emailTo = "N/A";
  if (req.user.managerId) {
    const manager = users.find((u) => u.id === req.user.managerId);
    emailTo = manager ? manager.username : "N/A";
  }

  app.locals.lastEmail = {
    to: emailTo,
    subject: `Approve expense ${id}`,
    body: `Approve here: /approve?token=${approvalToken}`,
  };

  res.json({
    expenseId: id,
    message: "Expense created and manager notified (simulated).",
  });
});

// View the "simulated" last email - available to any authenticated user in this challenge.
// Developer accidentally left this endpoint public to authenticated users.
app.get("/emails/last", auth, (req, res) => {
  if (!app.locals.lastEmail) return res.json({ message: "No emails yet" });
  res.json(app.locals.lastEmail);
});

// Manager approves using token link
app.get("/approve", (req, res) => {
  const { token } = req.query;
  if (!token) return res.status(400).send("missing token");

  try {
    const payload = jwt.verify(token, JWT_SECRET);

    // Business rule: only manager can approve, so we check managerId matches authenticated user.
    // But: the dev used X-User header to identify user, therefore if an attacker sets X-User to manager id,
    // they can "act" as the manager. The dev assumed the token was secret; it's viewable in /emails/last.
    const manager = users.find((u) => u.id === payload.managerId);
    if (!manager) return res.status(400).send("invalid manager in token");

    // The dev didn't require the requester to actually be the manager server-side
    // — they only check payload.managerId === payload.managerId (nonsense). This is the logic bug.
    const expense = expenses.find((e) => e.id === payload.expenseId);
    if (!expense) return res.status(404).send("expense not found");

    expense.status = "approved";
    expense.approvedBy = manager.username;

    return res.send(`Expense ${expense.id} approved by ${manager.username}`);
  } catch (err) {
    return res.status(400).send("invalid token");
  }
});

// Finance marks expense as paid (protected)
app.post("/pay/:id", auth, (req, res) => {
  if (req.user.role !== "finance")
    return res.status(403).json({ error: "only finance can pay" });

  const id = parseInt(req.params.id, 10);
  const expense = expenses.find((e) => e.id === id);
  if (!expense) return res.status(404).json({ error: "expense not found" });

  if (expense.status !== "approved")
    return res.status(400).json({ error: "expense not approved" });

  expense.status = "paid";

  // When an expense transitions to 'paid', the system writes the FLAG to a file readable only by finance.
  // In the challenge, the "flag" is revealed via the /flag route only if a paid expense exists.
  app.locals.paidExpense = expense.id;

  res.json({ message: `Expense ${id} paid.` });
});

// Endpoint to fetch the flag — only practical once a paid expense exists.
app.get("/flag", auth, (req, res) => {
  // Only pretend-protect: in challenge, if a paid expense exists, anyone can fetch /flag
  // to see "the secret" — the goal is to cause the system to reach that paid state.
  if (!app.locals.paidExpense)
    return res.status(404).json({ error: "No paid expenses yet" });
  const flagPath = path.join(__dirname, "FLAG");
  if (!fs.existsSync(flagPath))
    return res.status(500).json({ error: "FLAG file missing on server" });
  const flag = fs.readFileSync(flagPath, "utf8");
  res.type("text").send(flag);
});

// Serve embedded single-file HTML frontend
app.get("/", (req, res) => {
  res.type("html").send(`<!doctype html>
<html>
<head>
  <meta charset="utf-8" />
  <title>Approve Now — BizLogic Bypass</title>
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <style>
    body { font-family: system-ui, -apple-system, "Segoe UI", Roboto, "Helvetica Neue", Arial; background:#f7fafc; color:#111827; padding:20px; }
    .container { max-width:1000px; margin:0 auto; }
    header { display:flex; justify-content:space-between; align-items:center; margin-bottom:18px; }
    .card { background:white; border-radius:8px; padding:16px; box-shadow:0 1px 3px rgba(0,0,0,0.08); margin-bottom:12px; }
    .controls { display:flex; gap:8px; flex-wrap:wrap; }
    input, select, button, textarea { font:inherit; padding:8px; border:1px solid #e5e7eb; border-radius:6px; }
    button { cursor:pointer; background:#2563eb; color:white; border:none; }
    button.secondary { background:#6b7280; }
    pre { background:#0f172a; color:#d1fae5; padding:10px; border-radius:6px; overflow:auto; max-height:260px; }
    .muted { color:#6b7280; font-size:13px; }
    .small { font-size:13px; }
  </style>
</head>
<body>
  <div class="container">
    <header>
      <div>
        <h1>Approve Now — BizLogic Bypass</h1>
         
      </div>
      <div>
        <label class="small">User</label>
        <select id="actor">
          <option value="1">alice (employee)</option>
        </select>
        <button id="btn-login" class="secondary" style="margin-left:8px;">Login</button>
      </div>
    </header>

    <section class="card">
      <h3>Submit Expense</h3>
      <div style="display:flex; gap:8px; align-items:center; margin-bottom:8px;">
        <label class="small">Amount</label>
        <input id="amount" type="number" value="100" />
        <label class="small">Description</label>
        <input id="description" type="text" value="business travel" style="flex:1;" />
        <div class="controls">
          <button id="btn-submit">Submit</button>
          <button id="btn-view-email" class="secondary">View Last Email</button>
        </div>
      </div>
      <div id="email-box" class="small muted">No email fetched yet.</div>
    </section>

    <section class="card">
      <h3>Actions</h3>
      <div style="display:flex; gap:8px; align-items:center;">
        <input id="token-input" type="text" placeholder="Paste approval token or auto-fill from email" style="flex:1;" />
        <button id="btn-approve">Approve (use token)</button>
        <button id="btn-pay" class="secondary">Pay (as current actor)</button>
        <button id="btn-flag" style="background:#059669">Fetch Flag</button>
      </div>
    </section>

    <footer style="margin-top:12px;" class="muted small">
      <div>BizLogic Bypass CTF Challenge Demo - Developer made intentional security mistakes for learning purposes.</div>
    </footer>
  </div>

  <<script>
  const actorEl = document.getElementById('actor');
  const amountEl = document.getElementById('amount');
  const descEl = document.getElementById('description');
  const emailBox = document.getElementById('email-box');
  const tokenInput = document.getElementById('token-input');

  function addLog(line) {
    // now just log to console
    console.log(new Date().toISOString() + '  ' + line);
  }

  async function api(path, opts = {}) {
    opts.headers = opts.headers || {};
    opts.headers['X-User'] = actorEl.value;
    if (opts.body && typeof opts.body === 'object') {
      opts.body = JSON.stringify(opts.body);
      opts.headers['Content-Type'] = 'application/json';
    }
    try {
      const res = await fetch(path, opts);
      const text = await res.text();
      let json = null;
      try { json = JSON.parse(text); } catch(e) {}
      return { ok: res.ok, status: res.status, text, json };
    } catch (err) {
      return { ok: false, status: 0, text: String(err) };
    }
  }

  document.getElementById('btn-login').addEventListener('click', async () => {
    const username = actorEl.value === '1' ? 'alice' : actorEl.value === '2' ? 'bob' : 'charlie';
    addLog('Logging in as ' + username);
    const r = await api('/login', { method: 'POST', body: { username } });
    addLog('Login response: ' + (r.json ? JSON.stringify(r.json) : r.text));
  });

  document.getElementById('btn-submit').addEventListener('click', async () => {
    const amount = Number(amountEl.value) || 0;
    const description = descEl.value || '';
    addLog('Submitting expense: ' + amount + ' - ' + description);
    const r = await api('/expenses', { method: 'POST', body: { amount, description } });
    addLog('Submit response (' + r.status + '): ' + (r.json ? JSON.stringify(r.json) : r.text));
  });

  document.getElementById('btn-view-email').addEventListener('click', async () => {
    addLog('Fetching last email');
    const r = await api('/emails/last', { method: 'GET' });
    addLog('Email response (' + r.status + '): ' + (r.json ? JSON.stringify(r.json) : r.text));
    if (r.json && r.json.body) {
      emailBox.textContent = 'To: ' + r.json.to + ' | Subject: ' + r.json.subject + ' | Body: ' + r.json.body;
      const m = (r.json.body || '').match(/token=([A-Za-z0-9-_\\.]+)/);
      if (m) tokenInput.value = m[1];
    } else {
      emailBox.textContent = 'No email.';
    }
  });

  document.getElementById('btn-approve').addEventListener('click', async () => {
    const token = tokenInput.value.trim();
    if (!token) { addLog('No token provided'); return; }
    addLog('Calling /approve with token (as X-User=' + actorEl.value + ')');
    const r = await api('/approve?token=' + encodeURIComponent(token), { method: 'GET' });
    addLog('Approve response (' + r.status + '): ' + (r.text || ''));
  });

  document.getElementById('btn-pay').addEventListener('click', async () => {
    addLog('Attempting to pay latest expense (as X-User=' + actorEl.value + ')');
    let tryId = 1;
    const emailRes = await api('/emails/last', { method: 'GET' });
    if (emailRes.json && emailRes.json.subject) {
      const m = (emailRes.json.subject || '').match(/Approve expense (\d+)/);
      if (m) tryId = Number(m[1]);
    }
    const r = await api('/pay/' + tryId, { method: 'POST' });
    addLog('Pay response (' + r.status + '): ' + (r.json ? JSON.stringify(r.json) : r.text));
  });

  document.getElementById('btn-flag').addEventListener('click', async () => {
    addLog('Fetching /flag (as X-User=' + actorEl.value + ')');
    const r = await api('/flag', { method: 'GET' });
    addLog('Flag response (' + r.status + '): ' + (r.text || ''));
    if (r.ok) alert('Flag: ' + r.text);
  });

  addLog('UI ready. Select actor and interact with the server.');
</script>

</body>
</html>`);
});

// start server
app.listen(PORT, () => {
  console.log(`BizLogic Bypass listening on port ${PORT}`);
});
