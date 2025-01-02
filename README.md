
# Proof Of Work Blockchain

Proof of work Blockchain, finds a nonce for a valid block to add to the blockchain.


  - Blocks store data, position, Timestamp, hash, previousHash.
  - Each blocks previous Hash is linked to the one before it and that create blockchain.


## Create Folder

`mkdir project`

`cd project`

`go mod tidy`

`go run main.go`

## Port 

`http://localhost:8080`

## Test Endpoint

```
  curl -X POST http://localhost:8080 -H "Content-Type: application/json" -d '{"name": "89"}'
```

### Nonce
<img width="669" alt="Screenshot 2025-01-02 at 8 31 04 PM" src="https://github.com/user-attachments/assets/577d75c5-377e-4b23-9eec-8a6c161375df" />

### Output
<img width="703" alt="Screenshot 2025-01-02 at 8 27 38 PM" src="https://github.com/user-attachments/assets/3c86fbec-e113-40e5-8164-9cefa6448e1a" />


<img width="730" alt="Screenshot 2025-01-02 at 8 29 03 PM" src="https://github.com/user-attachments/assets/0125e01e-2be6-45e6-81a9-a3e5aa21c0b3" />




