# Wisdom - A Word of Wisdom TCP Server

The Wisdom project is a simple and secure TCP server that sends wisdom quotes to connected clients. To protect against DDoS attacks, the server utilizes a Proof-of-Work (PoW) mechanism using SHA-256 for client validation.

## Technologies Used

- Go
- Docker

### Algorithm: SHA-256

The SHA-256 algorithm is famously used in Bitcoin's PoW mechanism. Its capability to be 
both secure and adjustable makes it a time-tested choice for PoW algorithms.
The server uses the SHA-256 algorithm for PoW validation. We chose SHA-256 for its 
versatility and adaptability for various use-cases:

- **Simplicity**: The straightforwardness of the SHA-256 algorithm makes it an excellent 
choice for simpler devices. Its implementation is not overly complicated, making it 
accessible for a wide range of hardware.

- **Configurable Difficulty**: One of the primary advantages of using SHA-256 is its 
adjustable difficulty level. By modifying the number of leading zeros required in the 
hash, you can control how long it takes to solve the challenge. This adaptability
allows you to set the time for solving a single request, which is particularly 
useful for mitigating DDoS attacks without overwhelming genuine clients.

## Server Configuration

### Difficulty Level

The `--difficulty` parameter allows you to set the level of computational difficulty for the PoW validation **from the server-side**. Clients do not need to adjust this parameter; it's entirely controlled by the server. Recommended values are between 5-8, and the default value is 5.

### Run with Docker

To build and run the server using Docker, execute the following command:

```bash
make run-server
```

To build and run the client using Docker, execute the following command:
```bash
make run-client
```


### How to run
Server `go run cmd/server/server.go --network=tcp --address=0.0.0.0:8080 --difficulty=5`

Client `go run cmd/client/client.go --network=tcp --address=0.0.0.0:8080`

###
Dificulty can be configured with param `difficulty`.

Recommended value 5-8.

Default value: 4


### Further improvements
- Use logger Zap/Logrus
- Consider adding memory consuming algorithms to avoid attacks from GPU f.e. Scrypt or Argon2