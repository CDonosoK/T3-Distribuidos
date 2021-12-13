package main

import (

	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"github.com/CDonosoK/T3-Distribuidos/chat"

)

func crearCarpeta(directorio string){
	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.MkdirAll(directorio, 0755)
		if err != nil {
			log.Fatalf("Error creando la carpeta: %v", err)
		}
	}
}

func main(){
	//Conexión informantes - fulcrum 1
	crearCarpeta("Logs")
	crearCarpeta("Registros Planetarios")

	lis2, err2 := net.Listen("tcp", ":9002")
	if err2 != nil {
		log.Fatalf("Failed to listen on port 9002: %v", err2)
	}

	s2 := chat.Server{}

	grpcServer2 := grpc.NewServer()

	chat.RegisterChatServer(grpcServer2, &s2)

	if err2 := grpcServer2.Serve(lis2); err2 != nil {
		log.Fatalf("Failed to serve gRPC server over port 9002: %v", err2)

	}

}