# Build the manager binary
FROM golang:1.18 as builder

# ENV GO111MODULE=on GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
WORKDIR /workspace
# Copy the Go Modules manifests
COPY ./go.mod go.mod
COPY ./go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY ./main.go main.go
COPY ./client client
COPY ./grpc grpc

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o xca-api main.go

FROM alpine:latest as prod
WORKDIR /
COPY --from=builder /workspace/xca-api .
USER 8000:8000

ENTRYPOINT ["/cxa-api"]
