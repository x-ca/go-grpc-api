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
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"

	xgrpc "github.com/x-ca/go-grpc-api/grpc"
)

func TestClient(t *testing.T) {
	rpcClient, ctx, err := Client("127.0.0.1:8000")
	if err != nil {
		t.Logf("Get client err %s, skip.", err.Error())
		return
	}

	// version
	version, err := rpcClient.Version(ctx, &emptypb.Empty{})
	if err != nil {
		t.Logf("error happen when call gRPC api Version: %s", err.Error())
		return
	}
	t.Logf("version: %s", version)

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
		t.Logf("error happen when call gRPC api Sign: %s", err.Error())
		return
	}
	t.Logf("result: %s", result)
}
