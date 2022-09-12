package client

import (
	"log"
	"testing"

	"google.golang.org/protobuf/types/known/emptypb"

	xgrpc "github.com/x-ca/tls-grpc-api/grpc"
)

func TestClient(t *testing.T) {
	rpcClient, ctx, err := Client("localhost:8000")
	if err != nil {
		t.Logf("Get client err %s, skip.", err.Error())
	}

	// version
	version, err := rpcClient.Version(ctx, &emptypb.Empty{})
	if err != nil {
		t.Fatalf("error happen when call gRPC client for Version: %s", err.Error())
	}
	log.Printf("version: %s", version)

	// sign
	req := xgrpc.TLSRequest{
		CN:      "xiexianbin.cn",
		Domains: []string{"xiexianbin.cn", "*.xiexianbin.cn"},
		IPs:     []string{},
		Days:    10,
		KeyBits: 512,
	}
	result, err := rpcClient.Sign(ctx, &req)
	if err != nil {
		t.Fatalf("error happen when call gRPC client Sign: %s", err.Error())
	}
	t.Logf("result: %s", result)
}
