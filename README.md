
# Hashout - Seamless Checkout

![](https://img.shields.io/badge/coverage-95.8%25-brightgreen) ![](https://img.shields.io/github/go-mod/go-version/vinitius/hashout)

This production-ready service exposes a simple Checkout/Cart Rest API.

It also calculates discounts based off an external [discount's gRPC service](https://hub.docker.com/r/hashorg/hash-mock-discount-service).

  

# Arch

  

# Docs


| Method       | Resource            | Description                             |
|--------------|---------------------|-----------------------------------------|
| POST         |[/checkout]()        | Checkout by a set of pre-defined rules. |
| GET          |[/ping]()            | Healthcheck.                            |
| GET          |[/docs/index.html]() | Swagger-UI with an API Playground.      |

  

Import the [API specs]() into your favorite Rest Client.

 Generate docs:

```
make docs
```

  

# Dependencies

- `go` >= 1.16

- `make` (optional for better build experiences)

- `Docker/Docker-Compose` (optional for better deploy experiences)

  

# Main Dev Dependencies

- `Wire` (DI)

- `Gin` (http)

- `Swaggo` (swagger)

- `Testify/Mockery`(unit tests/mocks)

  

# Build

Download deps:

```
make deps
```

Go build:
```
make build
```

Dependency Injection:
```
make di
```

If you feel more comfortable running `go` commands directly, take a look at the [Makefile]() to check what you need.
  

# Test

Generate mocks:

```
make mocks
```

Run tests and generate coverage report:

```
make tests
```

  

# Run

You can either:

```
docker-compose up -d
```

Or:

```
make run
```

Or even:

```
go run cmd/main.go
```

Also:

Run in your favorite IDE or just run the single container instead of compose.

  

# Proto

You need [protoc]() and then:

```
make proto
```

# Profiling
25000 Requests: `Ubuntu x64 i7 32GB`
```
Summary:
  Total:	2.2339 secs
  Slowest:	0.0510 secs
  Fastest:	0.0002 secs
  Average:	0.0043 secs
  Requests/sec:	11191.1314
  
  Total data:	4625000 bytes
  Size/request:	185 bytes

Response time histogram:
  0.000 [1]	|
  0.005 [18417]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [3392]	|■■■■■■■
  0.015 [1717]	|■■■■
  0.020 [725]	|■■
  0.026 [400]	|■
  0.031 [160]	|
  0.036 [82]	|
  0.041 [33]	|
  0.046 [51]	|
  0.051 [22]	|


Latency distribution:
  10% in 0.0003 secs
  25% in 0.0006 secs
  50% in 0.0018 secs
  75% in 0.0056 secs
  90% in 0.0121 secs
  95% in 0.0165 secs
  99% in 0.0281 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0002 secs, 0.0510 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0028 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0057 secs
  resp wait:	0.0043 secs, 0.0001 secs, 0.0509 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0046 secs

Status code distribution:
  [200]	25000 responses
```

  

# SRE
 - Simple Healthcheck
 - Deploy runbook

  

# TBD

- Add cache layer with [Redis]()

- Improve validation messages

- Add integration tests suite with [TestContainers]()

- Add lint

- Buffer discount requests

- SRE: rollback runbook, metrics