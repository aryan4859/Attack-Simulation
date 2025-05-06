class LCG:
    def __init__(self, seed, a, c, m):
        self.state = seed
        self.a = a
        self.c = c
        self.m = m

    def next(self):
        self.state = (self.a * self.state + self.c) % self.m
        return self.state

class LCGCryptography:
    def __init__(self):
        # Hardcode LCG parameters (no secret key needed)
        self.seed = 12345678  # Fixed seed
        self.a = 1103515245   # Commonly used multiplier
        self.c = 12345        # Commonly used increment
        self.m = 2**31        # Modulus (2^31 is common for LCGs)
        self.lcg = LCG(self.seed, self.a, self.c, self.m)

    def encrypt(self, plaintext):
        ciphertext = []
        for char in plaintext:
            key = self.lcg.next() % 256
            encrypted_char = chr((ord(char) + key) % 256)
            ciphertext.append(encrypted_char)
        return ''.join(ciphertext)

    def decrypt(self, ciphertext):
        # Reset LCG to its initial state before decryption
        self.lcg = LCG(self.seed, self.a, self.c, self.m)
        plaintext = []
        for char in ciphertext:
            key = self.lcg.next() % 256
            decrypted_char = chr((ord(char) - key) % 256)
            plaintext.append(decrypted_char)
        return ''.join(plaintext)

# Example usage
crypto = LCGCryptography()

plaintext = "2ruqoM5RF89SzuZeGQEUw9D8owht1ykD"
ciphertext = crypto.encrypt(plaintext)
print(f"Ciphertext: {ciphertext}")

decrypted_text = crypto.decrypt(ciphertext)
print(f"Decrypted Text: {decrypted_text}")

# Verify decryption
assert decrypted_text == plaintext, "Decryption failed!"
print("Encryption and decryption process verified successfully!")