package client

import (
	"context"
	"fmt"
	"time"

	googlerpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	xgrpc "github.com/x-ca/tls-grpc-api/grpc"
)

// Client TLS GRPC API Client
// params target "127.0.0.1:8000"
func Client(target string) (xgrpc.ServiceClient, context.Context, error) {
	var creds credentials.TransportCredentials
	var err error
	creds = insecure.NewCredentials()
	cc, err := googlerpc.Dial(target, googlerpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, nil, err
	}
	defer func(cc *googlerpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(cc)

	rpcClient := xgrpc.NewServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return rpcClient, ctx, nil
}
