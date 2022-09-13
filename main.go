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

package main

import (
	"flag"
	"log"
	"net"

	googlerpc "google.golang.org/grpc"

	xgrpc "github.com/x-ca/go-grpc-api/grpc"
	xgrpcserver "github.com/x-ca/go-grpc-api/grpc/server"
)

var TLSServer = &xgrpcserver.TLSServiceServer{}
var (
	h              bool
	TLSKeyPath     string
	TLSCertPath    string
	TLSKeyPassword string
)

func init() {
	flag.BoolVar(&h, "help", false, "show help message")
	flag.StringVar(&TLSKeyPath, "tls-key", "", "tls key file path")
	flag.StringVar(&TLSCertPath, "tls-crt", "", "tls crt file path")
	flag.StringVar(&TLSKeyPassword, "tls-password", "", "tls key password")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	if TLSKeyPath == "" || TLSCertPath == "" {
		log.Fatal("TLSKeyPath or TLSCertPath is empty, start fail...")
	}

	_ = xgrpcserver.InitConfig(TLSKeyPath, TLSCertPath, TLSKeyPassword)
}

func main() {
	if h == true {
		flag.Usage()
		return
	}

	// Listener
	addr := ":8000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Println("grpc server closed.")
	}(listener)

	var s *googlerpc.Server
	s = googlerpc.NewServer()
	xgrpc.RegisterServiceServer(s, TLSServer)
	log.Println("grpc server listen on", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
