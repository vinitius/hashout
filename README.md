
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

  

Import the [API specs](handlers/server/swagger.json) into your favorite Rest Client.

 Generate docs:

```
make docs
```

  

# Dependencies

- `go` >= 1.16

- `make` (optional for a better build experience)

- `Docker/Docker-Compose` (optional for a better deploy experience)

  

# Main Dev Dependencies

- `Wire` (DI)

- `Gin` (http)

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

If you feel more comfortable running `go` commands directly, take a look at the [Makefile](Makefile) to check what you need.
  

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

You need [protoc](https://developers.google.com/protocol-buffers/docs/gotutorial) and then:

```
make proto
```

# Profiling
 `Ubuntu x64 i7 32GB`

*WEB:*
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

*CPU:*
```
Showing nodes accounting for 4.01s, 79.56% of 5.04s total
      flat  flat%   sum%        cum   cum%
     0.02s   0.4% 66.67%      4.11s 81.55%  net/http.(*conn).serve
     0.01s   0.2% 79.17%         3s 59.52%  viniti.us/hashout/handlers.CheckoutHandler.checkout
     0.01s   0.2% 79.37%      0.04s  0.79%  viniti.us/hashout/usecase/checkout.UseCase.IsBlackFridayGiftActive
     0.01s   0.2% 79.56%      1.87s 37.10%  viniti.us/hashout/usecase/discounts.UseCase.CalculateDiscounts
```

*MEM:*
```
Showing nodes accounting for 9911.22kB, 97.37% of 10178.83kB total
      flat  flat%   sum%        cum   cum%
   9760kB   95.89% 95.89%     9760kB 95.89%  golang.org/x/net/webdav.(*memFile).Write
   63.38kB  0.62%  97.29%    63.38kB  0.62%  bufio.NewWriterSize (inline)
   8.22kB   0.081% 97.37%    54.51kB  0.54%  net/http.(*conn).readRequest
         0     0% 97.37%   125.98kB  1.24%  viniti.us/hashout/handlers.CheckoutHandler.checkout
         0     0% 97.37%   125.98kB  1.24%  viniti.us/hashout/handlers.apiErrorReporter.func1
         0     0% 97.37%   113.77kB  1.12%  viniti.us/hashout/usecase/checkout.UseCase.Checkout
```

  

# SRE
 - Simple Healthcheck
 - Deploy runbook

  

# TBD

- Add cache layer with [Redis](https://redis.io/)

- Improve validation messages

- Add integration tests suite with [TestContainers](https://www.testcontainers.org/)

- Add lint

- Buffer discount requests

- SRE: rollback runbook, metrics