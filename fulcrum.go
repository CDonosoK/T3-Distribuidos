package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"

	"github.com/CDonosoK/T3-Distribuidos/chat"
)

type infoPlaneta struct {
	NomPlaneta string
	X          int32
	Y          int32
	Z          int32
}

type reloj struct{
	x int32
	y int32
	z int32
}

var relojServidor reloj
relojServidor.x = 0
relojServidor.y = 0
relojServidor.z = 0


func (s *Server) AddCityF(ctx context.Context, message *Message) (*Message, error) {
	//agregar la ciudad al archivo de registros
	planeta := message.Planeta
	//registros
	var ruta = "Registros/"+planeta+".txt"
	_, err := os.Stat(ruta)
	if os.IsNotExist(err){
		var file, err = os.Create(ruta)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		infoPlaneta = append(infoPlaneta, infoPlaneta{planeta, 0,0,0})
	}

	//logs
	var rutaL = "Logs/"+planeta+".txt"
	_,err := os.Stat(rutaL)
	if os.IsNotExist(err){
		var file2, err := os.Create(rutaL)
		if err != nil{
			panic(err)
		}
		defer file2.Close()
	}
	//reloj
	if message.Servidor == 0 {
		relojServidor.x += 1
	}
	if message.Servidor == 1 {
		relojServidor.y += 1
	}
	if message.Servidor == 2 {
		relojServidor.z += 1
	}
	
	//cambiar el return
	
	return &Message{Servidor: server}, nil
}

func (s *Server) UpdateNameF(ctx context.Context, message *Message) (*Message, error) {
	
	planeta = message.Planeta
	ciudad = message.Ciudad
	valor = message.Valor
	var ruta = "Registros/"+planeta+".txt"
	file, err := ioutil.ReadFile(ruta)
	linea := strings.Split(string(file), " ")

	nuevoCiudad := linea[1]
	lineaN = planeta + " "+nuevoCiudad+" "+valor+"\n"
	//escribir archivo de registro nuevo
	err = ioutil.WriteFile(ruta, []byte(lineaN), 0664)
	/*if err != {
		log.Fatal(err)
		//mostrar algun mensaje de error?
	}*/
	//logs
	log_m := "UpdateName "+planeta+" "+ciudad+" "+nuevoCiudad+"\n"
	file,err := ioutil.ReadFile(ruta)
	file = file.append(file, []byte(log_m)...)
	err = ioutil.WriteFile(rutaL, file, 0644)
		
	//reloj
	if message.Servidor == 0 {
		relojServidor.x += 1
	}
	if message.Servidor == 1 {
		relojServidor.y += 1
	}
	if message.Servidor == 2 {
		relojServidor.z += 1
	}

	//log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Servidor: serverElegido}, nil
}


func (s *Server) UpdateNumberF(ctx context.Context, message *Message) (*Message, error) {
	valor = message.Valor
	var ruta = "Registros/"+planeta+".txt"
	file, err := ioutil.ReadFile(ruta)
	linea := strings.Split(string(file), " ")

	nuevoValor := linea[2]
	lineaN = nuevoNombre + " "+ciudad+" "+valor+"\n"
	//escribir archivo de registro nuevo
	err = ioutil.WriteFile(ruta, []byte(lineaN), 0664)
	if err != {
		log.Fatal(err)
		//mostrar algun mensaje de error?
	}
	//logs
	log_m := "UpdateNumber "+planeta+" "+valor+" "+nuevoValor+"\n"
	file,err := ioutil.ReadFile(ruta)
	file = file.append(file, []byte(log_m)...)
	err = ioutil.WriteFile(rutaL, file, 0644)
		
	//reloj
	if message.Servidor == 0 {
		relojServidor.x += 1
	}
	if message.Servidor == 1 {
		relojServidor.y += 1
	}
	if message.Servidor == 2 {
		relojServidor.z += 1
	}

	//log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Servidor: serverElegido}, nil
}

func (s *Server) DeleteCityF(ctx context.Context, message *Message) (*Message, error) {
	var ruta = "Registros/"+planeta+".txt"
	file, err := ioutil.ReadFile(ruta)
	linea := strings.Split(string(file), " ")
	planeta = linea[0]
	ciudad = linea[1]
	valor = linea[2]
	lineaN = planeta +" "+valor
	//escribir en registro
	err = ioutil.WriteFile(ruta, []byte(lineaN), 0664)
	if err != {
		log.Fatal(err)
		//mostrar algun mensaje de error?
	}
	//logs
	log_m := "DeleteCity "+planeta+" "+ciudad+"\n"
	file,err := ioutil.ReadFile(ruta)
	file = file.append(file, []byte(log_m)...)
	err = ioutil.WriteFile(rutaL, file, 0644)

	//reloj
	if message.Servidor == 0 {
		relojServidor.x += 1
	}
	if message.Servidor == 1 {
		relojServidor.y += 1
	}
	if message.Servidor == 2 {
		relojServidor.z += 1
	}


	//log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Servidor: serverElegido}, nil
}

func main() {
	registros := filepath.Join(".", "RegistrosPlanetarios")
	logs := filepath.Join(".", "logs")

	//verifica que la carpeta no exista
	if _, err := os.Stat(registros); os.IsNotExist(err) {
		err := os.Mkdir(registros, 0755)
		if err != nil {
			log.Fatalf("Fallo en crear la carpeta: %v", err)
		}
	}
	if _, err := os.Stat(logs); os.IsNotExist(err) {
		err := os.Mkdir(logs, 0755)
		if err != nil {
			log.Fatalf("Fallo en crear la carpeta: %v", err)
		}
	}

	/*// colocar servidor correcto
	lis2, err2 := net.Listen("tcp", ":9002")
	if err2 != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err2)
	}
	defer lis2.Close()

	s0 := chat.Server{}

	grpcServer0 := grpc.NewServer()

	chat.RegisterChatServer(grpcServer0, &s0)
	fmt.Println("Servidor en 9002")

	if err2 := grpcServer0.Serve(lis2); err2 != nil {
		log.Fatalf("Failed to serve gRPC server over port 8000: %v", err2)
	}*/
}
