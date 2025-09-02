Solution

First let es be e i and cs be c i .

Suppose you found a i such that ∑ i = 0 n − 1 a i e i = 0 , then the following equation will be true:

∏ i = 0 n − 1 c i a i = m e i a i = m 0 = 1 ( mod n )

So finding 2 such a i and we can recover n with gcd, but since we are computing the power in ring of integers so we need to make them small. And LLL can be used here to find some small linear combinations for that:

L = [ K e 0 1 K e 1 1 ⋮ ⋱ K e n − 1 1 ]

The first column of first few vectors in the reduced basis of L will be 0 when K is large enough, and the remaining columns will captures those a i for us.

Also, a i will have some negative entries in it, so ∏ i = 0 n − 1 c i a i will be a rational number. Denote it as x y then we will have this:

x y ≡ 1 ( mod n ) x y − 1 ≡ 0 ( mod n ) x − y ≡ 0 ( mod n )

So gcd ( x 0 − y 0 , x 1 − y 1 ) will probably be n .

After getting n , the flag can be easily decrypted with common modulus attack.