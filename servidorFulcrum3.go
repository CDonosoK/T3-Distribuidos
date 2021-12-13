package main

import (

	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/CDonosoK/T3-Distribuidos/chat"

)

func main(){
	//Conexión informantes - fulcrum 3
	lis4, err4 := net.Listen("tcp", ":9004")
	if err4 != nil {
		log.Fatalf("Failed to listen on port 9004: %v", err4)
	}

	s4 := chat.Server{}

	grpcServer3 := grpc.NewServer()

	chat.RegisterChatServer(grpcServer3, &s4)

	if err4 := grpcServer3.Serve(lis4); err4 != nil {
		log.Fatalf("Failed to serve gRPC server over port 9004: %v", err4)

	}

}