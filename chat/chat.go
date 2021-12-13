package chat

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"golang.org/x/net/context"
)

type Server struct {

}

type infoPlaneta struct {
	NomPlaneta string
	ciudad 	string
	valor string
}

type relojPlaneta struct {
	X int32
	Y int32
	Z int32
}

var listaPlaneta []infoPlaneta
var reloj = relojPlaneta{
	X: 0,
	Y: 0,
	Z: 0,
}

func (s *Server) AddCityF(ctx context.Context, message *Message) (*Message, error) {
	// SE OBTIENE LA INFORMACIÓN
	planeta := message.Planeta

	nuevoPlaneta := infoPlaneta{
		NomPlaneta: message.Planeta,
		ciudad: message.Ciudad,
		valor: message.Valor,
	}


	// SE ACTUALIZA EL RELOJ
	if message.Servidor == 0 {
		reloj.X += 1
	}
	if message.Servidor == 1 {
		reloj.Y += 1
	}
	if message.Servidor == 2 {
		reloj.Z += 1
	}
	

	listaPlaneta = append(listaPlaneta, nuevoPlaneta)
	fmt.Println(reloj.X, reloj.Y, reloj.Z)

	log.Printf("Planeta: %s", planeta)

	// SE CREAN LAS RUTAS CORRESPONDIENTES
	directorio1 := "./Logs/" + planeta + ".txt"
	directorio2 := "./Registros Planetarios/" + planeta + ".txt"

	// SE ESCRIBE EN EL LOGS - IGUAL PARA TODAS LAS FUNCIONES
	logMessage := "AddCity " + message.Planeta + " " + message.Ciudad + " " + message.Valor + "\n"
	f, err := os.OpenFile(directorio1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(logMessage); err != nil {
		log.Println(err)
	}

	// SE ESCRIBE EN EL REGISTRO
	registroMessage := message.Planeta + " " + message.Ciudad + " " + message.Valor + "\n"
	f1, err1 := os.OpenFile(directorio2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		log.Println(err1)
	}
	defer f1.Close()
	if _, err := f1.WriteString(registroMessage); err != nil {
		log.Println(err)
	}

	// Actualizar el return

	return &Message{Planeta: "RECIBIDO"}, nil
}

func (s *Server) AddCityMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))	

	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Valor: message.Valor, Servidor: serverElegido}, nil
}

func (s *Server) UpdateNameF(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Valor: message.Valor, Servidor: serverElegido}, nil
}

func (s *Server) UpdateNameMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Valor: message.Valor, Servidor: serverElegido}, nil
}

func (s *Server) UpdateNumberF(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Servidor: serverElegido}, nil
}

func (s *Server) UpdateNumberMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Valor: message.Valor, Servidor: serverElegido}, nil
}

func (s *Server) DeleteCityF(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Servidor: serverElegido}, nil
}

func (s *Server) DeleteCityMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Servidor: serverElegido}, nil
}

func (s *Server) ObtenerNumeroRebeldesBroker(ctx context.Context, message *DeLeia) (*ParaLeia, error) {


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n", message.Planeta, message.Ciudad)
	return &ParaLeia{X: 1}, nil
}

func (s *Server) ObtenerNumeroRebeldesFulcrum(ctx context.Context, message *DeLeia) (*ParaLeia, error) {


	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n", message.Planeta, message.Ciudad)
	return &ParaLeia{X: 1}, nil
}

