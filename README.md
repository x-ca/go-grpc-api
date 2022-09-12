# tls-grpc-api

[![build-test](https://github.com/x-ca/tls-grpc-api/actions/workflows/workflow.yaml/badge.svg)](https://github.com/x-ca/tls-grpc-api/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/x-ca/tls-grpc-api?status.svg)](https://pkg.go.dev/github.com/x-ca/tls-grpc-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/x-ca/tls-grpc-api)](https://goreportcard.com/report/github.com/x-ca/tls-grpc-api)

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

```
go run main.go \
  -tls-crt /Users/xiexianbin/workspace/code/github.com/kbcx/temp/x-ca/ca/tls-ca.crt \
  -tls-key /Users/xiexianbin/workspace/code/github.com/kbcx/temp/x-ca/ca/tls-ca/private/tls-ca.key
```
