package main

import (

	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/CDonosoK/T3-Distribuidos/chat"

)

func main(){
	//Conexión broker - informantes
	lis0, err0 := net.Listen("tcp", ":9000")
	if err0 != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err0)
	}

	s0 := chat.Server{}

	grpcServer0 := grpc.NewServer()

	chat.RegisterChatServer(grpcServer0, &s0)

	if err0 := grpcServer0.Serve(lis0); err0 != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err0)

	}

	//Conexión broker - leia
	lis1, err1 := net.Listen("tcp", ":9001")
	if err1 != nil {
		log.Fatalf("Failed to listen on port 9001: %v", err1)
	}

	s1 := chat.Server{}

	grpcServer1 := grpc.NewServer()

	chat.RegisterChatServer(grpcServer1, &s1)

	if err1 := grpcServer1.Serve(lis1); err1 != nil {
		log.Fatalf("Failed to serve gRPC server over port 8000: %v", err1)

	}

	//Conexión broker - servidores
	lis2, err2 := net.Listen("tcp", ":9002")
	if err2 != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err2)
	}

	s2 := chat.Server{}

	grpcServer2 := grpc.NewServer()

	chat.RegisterChatServer(grpcServer2, &s2)

	if err2 := grpcServer2.Serve(lis2); err2 != nil {
		log.Fatalf("Failed to serve gRPC server over port 8000: %v", err2)

	}

}