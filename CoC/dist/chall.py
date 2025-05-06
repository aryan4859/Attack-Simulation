import random

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
    def __init__(self, key):
        # Generate LCG parameters from the key
        random.seed(key)
        self.seed = random.randint(0, 2**32 - 1)
        self.a = random.randint(2, 2**16)
        self.c = random.randint(1, 2**16)
        self.m = 2**32
        self.lcg = LCG(self.seed, self.a, self.c, self.m)

    def encrypt(self, plaintext):
        ciphertext = []
        for char in plaintext:
            key = self.lcg.next() % 256
            encrypted_char = chr((ord(char) + key) % 256)
            ciphertext.append(encrypted_char)
        return ''.join(ciphertext)

    def decrypt(self, ciphertext):
        self.lcg = LCG(self.seed, self.a, self.c, self.m)  # Reset LCG to initial state
        plaintext = []
        for char in ciphertext:
            key = self.lcg.next() % 256
            decrypted_char = chr((ord(char) - key) % 256)
            plaintext.append(decrypted_char)
        return ''.join(plaintext)

# Example usage
key = "forge-that-flag"
crypto = LCGCryptography(key)

plaintext = "2ruqoM5RF89SzuZeGQEUw9D8owht1ykD"
ciphertext = crypto.encrypt(plaintext)
print(f"Ciphertext: {ciphertext}")

decrypted_text = crypto.decrypt(ciphertext)
print(f"Decrypted Text: {decrypted_text}")