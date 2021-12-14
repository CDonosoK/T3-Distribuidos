package main

import (

	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/CDonosoK/T3-Distribuidos/chat"

)

func main(){
	//Conexión broker - informantes - leia
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

	//Conexión informantes - fulcrum 1
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

	//Conexión informantes - fulcrum 2
	lis3, err3 := net.Listen("tcp", ":9003")
	if err3 != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err3)
	}

	s3 := chat.Server{}

	grpcServer3 := grpc.NewServer()

	chat.RegisterChatServer(grpcServer3, &s3)

	if err3 := grpcServer3.Serve(lis3); err3 != nil {
		log.Fatalf("Failed to serve gRPC server over port 8000: %v", err3)

	}

	//Conexión informantes - fulcrum 3
	lis4, err4 := net.Listen("tcp", ":9004")
	if err4 != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err4)
	}

	s4 := chat.Server{}

	grpcServer4 := grpc.NewServer()

	chat.RegisterChatServer(grpcServer4, &s4)

	if err4 := grpcServer4.Serve(lis4); err4 != nil {
		log.Fatalf("Failed to serve gRPC server over port 8000: %v", err4)

	}

}