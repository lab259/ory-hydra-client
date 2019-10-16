[![Actions Status](https://github.com/lab259/ory-hydra-client/workflows/Go/badge.svg)](https://github.com/lab259/ory-hydra-client/actions) [![GoDoc](https://godoc.org/github.com/lab259/ory-hydra-client?status.svg)](http://godoc.org/github.com/lab259/ory-hydra-client) [![Go Report Card](https://goreportcard.com/badge/github.com/lab259/ory-hydra-client)](https://goreportcard.com/report/github.com/lab259/ory-hydra-client) [![Coverage](https://gocover.io/_badge/github.com/lab259/ory-hydra-client)](http://gocover.io/github.com/lab259/ory-hydra-client)

# ory-hydra-client

## Getting Started

### Prerequisites

What things you need to setup the project:

- [go](https://golang.org/doc/install)
- [ginkgo](http://onsi.github.io/ginkgo/)

### Environment

Clone the repository:

```bash
git clone git@github.com:lab259/ory-hydra-client.git
```

Now, the dependencies must be installed.

```
cd ory-hydra-client && go mod download
```

:wink: Finally, you are done to start developing.

### Running tests

In the `src/github.com/lab259/ory-hydra-client` directory, execute:

```bash
make test
```

To enable coverage, execute:

```bash
make coverage
```

To generate the HTML coverage report, execute:

```bash
make coverage coverage-html
```
