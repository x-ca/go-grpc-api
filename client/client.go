// Copyright 2022 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"time"

	googlerpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	xgrpc "github.com/x-ca/go-grpc-api/grpc"
)

// Client TLS GRPC API Client
// params target "127.0.0.1:8000"
func Client(target string) (xgrpc.ServiceClient, context.Context, error) {
	var err error
	var creds credentials.TransportCredentials
	creds = insecure.NewCredentials()
	cc, err := googlerpc.Dial(target, googlerpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, nil, err
	}
	// use defer occur rpc error: code = Canceled desc = context canceled
	//defer func(cc *googlerpc.ClientConn) {
	//	err := cc.Close()
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//}(cc)

	rpcClient := xgrpc.NewServiceClient(cc)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	return rpcClient, ctx, nil
}
