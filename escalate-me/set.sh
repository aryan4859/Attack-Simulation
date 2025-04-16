#!/bin/bash

# Create CTF user
useradd -m -s /bin/bash ctfuser
echo "ctfuser:ctfuser" | chpasswd

# Create user-level flag
echo "FLAG{redacted}" > /home/ctfuser/user.txt
chmod 600 /home/ctfuser/user.txt
chown ctfuser:ctfuser /home/ctfuser/user.txt

# Create root-level flag
echo "FLAG{redacted}" > /root/root.txt
chmod 600 /root/root.txt
chown root:root /root/root.txt

# SUID binary for privilege escalation
cat <<EOF > /tmp/getroot.c
#include <stdlib.h>
int main() {
    setuid(0);
    system("/bin/bash");
    return 0;
}
EOF

gcc /tmp/getroot.c -o /usr/local/bin/getroot
chmod u+s /usr/local/bin/getroot
chown root:root /usr/local/bin/getroot
rm /tmp/getroot.c

# Clean bash history
history -c
unset HISTFILE
rm -f /home/ctfuser/.bash_history
rm -f /root/.bash_history

echo "[+] Setup complete. Login with ctfuser:ctfuser and find a way to read /root/root.txt"
