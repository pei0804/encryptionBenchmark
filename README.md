# encryptionBenchmark

```
$ glide install
$ go run main.go
```

## 1st time

- scrypt: 73.621115ms
- bcrypt: 52.998518ms
- sha256: 3.253µs
- sha512: 4.276µs
- sha3-256: 5.999µs
- Ripemd: 5.323µs
- MD5: 2.669µs

## 2nd time

- scrypt: 74.315378ms
- bcrypt: 51.524798ms
- sha256: 2.514µs
- sha512: 2.339µs
- sha3-256: 7.625µs
- Ripemd: 9.218µs
- MD5: 2.625µs

## 3rd time

- scrypt: 79.316009ms
- bcrypt: 50.530978ms
- sha256: 2.305µs
- sha512: 3.291µs
- sha3-256: 8.048µs
- Ripemd: 4.314µs
- MD5: 3.104µs

## 4th time

- scrypt: 73.314786ms
- bcrypt: 49.949264ms
- sha256: 2.296µs
- sha512: 3.083µs
- sha3-256: 7.453µs
- Ripemd: 8.307µs
- MD5: 1.775µs

## 5th time

- scrypt: 77.01785ms
- bcrypt: 50.199802ms
- sha256: 3.272µs
- sha512: 3.321µs
- sha3-256: 6.446µs
- Ripemd: 3.99µs
- MD5: 2.759µs
