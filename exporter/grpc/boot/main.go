/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"context"
	"gend/core"
	"gend/core/config"
	pb "gend/exporter/grpc"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{
	gendFacade core.GendFacade
}

// SayHello implements helloworld.GreeterServer
func (s *server) GenId(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{Id: s.gendFacade.GenId()}, nil
}
func (s *server) ExpId(ctx context.Context, in *pb.RequestLong) (*pb.ReplyString, error) {
	return &pb.ReplyString{Json: s.gendFacade.ExpId(in.Id).ToJson()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	config.InitConfig("")
	config.GetDefaultConfigInstance().Gen = 1
	gendFacade := core.GendFacade{}
	gendFacade.Init()
	gServer := server{gendFacade: gendFacade}
	s := grpc.NewServer()
	pb.RegisterGendRpcServer(s, &gServer)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}