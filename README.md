# word-of-wisdom-tcp-server

## Problem/Purpose

Design and implement “Word of Wisdom” tcp server.

- TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained.
- After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the POW challenge

## Why have I picked Hascash algorithm?

The primary goal of a PoW algorithm is to make computational work time-consuming and resource-intensive, thus deterring attackers from performing Distributed Denial of Service (DDoS) attacks.

### Hashcash

**Algorithm Overview**: Hashcash requires a party to perform a certain amount of computational work, which involves finding a partial hash collision. The work is typically to find a hash that starts with a specified number of leading zeros.

**Reasons for Choosing Hashcash**:

- **Widely Adopted**: Hashcash is a well-known and widely adopted PoW algorithm, used in various contexts, including email spam prevention.
- **Simple yet Effective**: The concept is straightforward, making it easy to implement and understand.
- **Resource Intensive**: The algorithm is computationally intensive, requiring significant CPU power to find a valid solution. This adds a layer of protection against DDoS attacks.
- **Proven Track Record**: Hashcash has been in use for years and has a proven track record of effectively mitigating certain types of attacks.

### Summary

Ultimately, the choice depends on the specific security requirements of the system, the desired level of complexity, and the trade-offs between computational resources and memory usage. In this case, **Hashcash is chosen for its simplicity, widespread use, and effectiveness in deterring DDoS attacks.**

## How to run a server?

Have a look into Dockerfile. It creates light-weight docker image on top of alpine where you have compiled binary `/app/word-of-wisdom-tcp-server`.

You can customize hostname and port if needed. Defaults: 0.0.0.0:8080

```shell
go run cmd/main.go -h

Usage: word-of-wisdom-tcp-server [OPTIONS]

Example: word-of-wisdom-tcp-server -p 8080

Options:
  -H, --hostname     listen on hostname (default "0.0.0.0")
  -p, --port         port to listen to (default "8080")
  -d, --difficulty   challenge difficulty (default "4")
```

### Run a server in Docker container

```shell
# Build docker image
$ docker build -t word-of-wisdom-server .

# Run an app in docker container
$ docker run -p 8080:8080 word-of-wisdom-server
# 2024/01/22 12:54:01 Server listening on 0.0.0.0:8080
```

### Run client

See [Client's README](https://github.com/serg-kovalev/word-of-wisdom-tcp-client/blob/main/README.md)

### Example

```shell
# simple difficulty of 4 leading zeros in a hash
$ go run cmd/main.go
2024/01/29 13:35:41 server listening on 0.0.0.0:8080
2024/01/29 13:35:51 challenge: 4:qpJhEIs9xuwFyoiRkoLkkgCjHboTX4rxZvdZ3ovrO9u9KdTqiz
2024/01/29 13:35:51 nonce:  4212
2024/01/29 13:35:51 calculated hash:  00007d4419c11cf8d788d10059d2cb17ea148b7c49fd137385b54e7726273085
2024/01/29 13:35:51 quote sent to client: Quote 8
2024/01/29 13:35:52 challenge: 4:bFsazympmX2VUO7gaC2ia0epXM93FvWOvvAswtgXDqFzA4ApJL
2024/01/29 13:35:52 nonce:  23157
2024/01/29 13:35:52 calculated hash:  00005697d474e50497329c2569b83940f84d091d953b7193fda6663a0e2ddc90
2024/01/29 13:35:52 quote sent to client: Quote 6
2024/01/29 13:35:55 challenge: 4:MhosKzSnx6sfEwT4wgEbQ6koh8W2NlSQ6BDsPi2ZPBMgeb2mp6
2024/01/29 13:35:55 nonce:  19350
2024/01/29 13:35:55 calculated hash:  000076df79048fd4cdf9162f602da7adef989a910f5e1acea52c510ecc436f4b
2024/01/29 13:35:55 quote sent to client: Quote 7

# difficulty 6:
$ go run cmd/main.go -d 6
2024/01/29 13:36:02 server listening on 0.0.0.0:8080
2024/01/29 13:36:07 challenge: 6:1Rk5hiaXsUfgKA5TJ7IBzdCTteYRqkyWHFTDBqwU4wvizo4XrT
2024/01/29 13:36:19 nonce:  41214698
2024/01/29 13:36:19 calculated hash:  00000045ce84396aa29ae3936e77fb7e2d1fb80f09a71124b9153a62397c157b
2024/01/29 13:36:19 quote sent to client: Quote 9
```
