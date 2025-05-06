# ⏱️ Timekeeper - CTF Challenge

**Category:** Web  
**Difficulty:** Medium  
**Vulnerability:** Username Enumeration via Response Timing  

---

## 📘 Description

Welcome to **Timekeeper** — a CTF challenge simulating a real-world vulnerability where the login response time leaks information about valid usernames.

The challenge features a login portal with a subtle timing difference when a correct username is provided versus an incorrect one.

---

## 🛠️ Setup Instructions

### 🐳 Using Docker (Recommended)

```bash
git clone https://github.com/your-repo/timekeeper-ctf.git
cd timekeeper-ctf
docker build -t timekeeper .
docker run -p 3000:3000 timekeeper
```
## Step-by-Step Exploitation
 - Step 1:Test Invalid Username Timing
 ```bash
curl -X POST -d "username=test&password=123" http://localhost:3000/login -w "\n%{time_total}\n"

 ```
 - Step 2: Test Valid Username Timing
```bash
curl -X POST -d "username=admin&password=123" http://localhost:3000/login -w "\n%{time_total}\n"

```

 - Step 3: 