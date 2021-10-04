
# Hashout - Seamless Checkout

![](https://img.shields.io/badge/coverage-95.8%25-brightgreen) ![](https://img.shields.io/github/go-mod/go-version/vinitius/hashout)

This production-ready service exposes a simple Checkout/Cart Rest API.

It also calculates discounts based off an external [discount's gRPC service](https://hub.docker.com/r/hashorg/hash-mock-discount-service).

  

# Arch

  

# Docs


| Method       | Resource            | Description                             |
|--------------|---------------------|-----------------------------------------|
| POST         |[/checkout]()        | Checkout by a set of pre-defined rules. |
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

You need protoc and then:

```
make proto
```

# Profiling

  

# SRE
 - Simple Healthcheck
 - Deploy runbook

  

# TBD

- Add cache layer with [Redis]()

- Improve validation messages

- Add integration tests suite with [TestContainers]()

- Add lint

- Buffer discount requests

- SRE: rollback runbook

- SRE: metrics