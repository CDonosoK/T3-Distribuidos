package main

import (

	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/CDonosoK/T3-Distribuidos/chat"
)

func main(){
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

}