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

package server

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/pem"
	"github.com/x-ca/go-ca/ca"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"strings"
	"sync"

	xgrpc "github.com/x-ca/go-grpc-api/grpc"
)

type Config map[string]string

var (
	once sync.Once
	conf Config
)

func InitConfig(keyPath, certPath, keyPassword string) Config {
	once.Do(func() {
		conf = make(Config)
		conf["TLSKeyPath"] = keyPath
		conf["TLSCertPath"] = certPath
		conf["TLSKeyPassword"] = keyPassword
	})
	return conf
}

type TLSServiceServer struct {
	xgrpc.UnimplementedServiceServer
}

func (s *TLSServiceServer) Sign(ctx context.Context, req *xgrpc.TLSRequest) (*xgrpc.TLSResponse, error) {
	log.Println("new sign request, req is", req.String())
	var err error

	tlsCA, err := ca.LoadTLSCA(conf["TLSKeyPath"], conf["TLSCertPath"], conf["TLSKeyPassword"])
	if err != nil {
		return nil, err
	}

	ips, err := ca.ParseIPs(req.IPs)
	if err != nil {
		return nil, err
	}

	keyBits := 1024
	if req.KeyBits%256 == 0 {
		keyBits = 1024
	}

	key, cert, err := tlsCA.Sign(req.CN, req.Domains, ips, int(req.Days), keyBits)
	if err != nil {
		return nil, err
	}

	keyBuf := new(strings.Builder)
	err = pem.Encode(keyBuf, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	certBuf := new(bytes.Buffer)
	err = pem.Encode(certBuf, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	resp := &xgrpc.TLSResponse{
		Key:  keyBuf.String(),
		Cert: certBuf.String(),
	}
	return resp, nil
}

func (s *TLSServiceServer) Version(ctx context.Context, empty *emptypb.Empty) (*xgrpc.VersionResponse, error) {
	version := &xgrpc.VersionResponse{
		Version: "v1.0.0",
	}
	return version, nil
}
