ciphertext = "¡îzËúµ¶x­Lö}õc¦ýº_òÑ5Æ»Ö$)ò"
# Take first N characters
N = 8
for i in range(N):
    c = ciphertext[i]
    ord_c = ord(c)
    print(f"\nChar {i}: {repr(c)} (ord: {ord_c})")
    # Try key guesses 0-255 and see printable results
    for k in range(256):
        p = (ord_c - k) % 256
        if 32 <= p <= 126:  # Printable ASCII
            print(f"  Key: {k:3} => Plain: {chr(p)}")
