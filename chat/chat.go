package chat

import (
	"log"
	"math/rand"
	"os"
	"fmt"

	"golang.org/x/net/context"

)

type Server struct {

}

type infoPlaneta struct {
	NomPlaneta string
	ciudad 	string
	valor string
	X          int32
	Y          int32
	Z          int32
}

var listaPlaneta []infoPlaneta

func (s *Server) AddCityF(ctx context.Context, message *Message) (*Message, error) {
	// Se obtiene la información
	planeta := message.Planeta

	nuevoPlaneta := infoPlaneta{
		NomPlaneta: message.Planeta,
		ciudad: message.Ciudad,
		valor: message.Valor,
		X: 0,
		Y: 0,
		Z: 0,
	}

	listaPlaneta = append(listaPlaneta, nuevoPlaneta)
	fmt.Println(listaPlaneta)

	log.Printf("Planeta: %s", planeta)

	// Se crean los archivos / registros
	directorio1 := "./Logs/" + planeta + ".txt"
	directorio2 := "./Registros Planetarios/" + planeta + ".txt"

	//SE ESCRIBE EN EL LOGS
	logMessage := "AddCity " + message.Planeta + " " + message.Ciudad + " " + message.Valor + "\n"
	f, err := os.OpenFile(directorio1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(logMessage); err != nil {
		log.Println(err)
	}

	//SE ESCRIBE EN EL REGISTRO
	//ESTO SE DEBERÍA MODIFICAR -> SE RECORRE EL STRUCT O LA LISTA
	registroMessage := message.Planeta + " " + message.Ciudad + " " + message.Valor + "\n"
	f1, err1 := os.OpenFile(directorio2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		log.Println(err1)
	}
	defer f1.Close()
	if _, err := f1.WriteString(registroMessage); err != nil {
		log.Println(err)
	}

	// Actualizar el reloj

	// Actualizar el return

	return &Message{Planeta: "RECIBIDO"}, nil
}

func (s *Server) AddCityMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))	

	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Servidor: serverElegido}, nil
}


func (s *Server) UpdateNameMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Servidor: serverElegido}, nil
}

func (s *Server) UpdateNumberMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Servidor: serverElegido}, nil
}

func (s *Server) DeleteCityMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Servidor: serverElegido}, nil
}

