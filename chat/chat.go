package chat

import(
	"log"
	"math/rand"

	"golang.org/x/net/context"
)

type Server struct {

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