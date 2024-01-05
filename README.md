## Transaction Broadcasting and Monitoring Client Module

go to `/server` folder and using

```
go run main.go
```

### Endpoints

1. **Broadcast Transaction**

- POST

  ```
  /broadcast
  ```

- BODY:
  ```json
  {
    "symbol": "string", // Transaction symbol, e.g., BTC
    "price": uint64, // Symbol price, e.g., 100000
    "timestamp": uint64 // Timestamp of price retrieval
  }
  ```
- Example:
  ```
  curl --location 'localhost:1323/mnc' \
  --header 'Content-Type: application/json' \
  --data '{
      "symbol": "ETH",
      "price": 45000,
      "timestamp": 167891234235
  }'
  ```

2. **Transaction Status Monitoring**

- GET
  ```
  /broadcast/<tx_hash>
  ```
- Example:
  ```
  curl --location 'localhost:1323/mnc/6774392a2a2e67efd2cd61b366dfcc2046bdec4ffc9a6f77ed98b88995f5f241'
  ```
