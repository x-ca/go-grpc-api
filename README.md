# X-CA go-grpc-api

[![build-test](https://github.com/x-ca/go-grpc-api/actions/workflows/workflow.yaml/badge.svg)](https://github.com/x-ca/go-grpc-api/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/x-ca/go-grpc-api?status.svg)](https://pkg.go.dev/github.com/x-ca/go-grpc-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/x-ca/go-grpc-api)](https://goreportcard.com/report/github.com/x-ca/go-grpc-api)

## Usage

```
$ go run main.go -h
  -help
    	show help message
  -tls-crt string
    	tls crt file path
  -tls-key string
    	tls key file path
  -tls-password string
    	tls key password
```

- [create ca file](https://github.com/x-ca/go-ca), start gRPC api

```
go run main.go \
  -tls-crt ./x-ca/ca/tls-ca.crt \
  -tls-key ./x-ca/ca/tls-ca/private/tls-ca.key
```
