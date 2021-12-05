package chat

import(
	"log"

	"golang.org/x/net/context"
)

type Server struct {

}

func (s *Server) AddCityMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Planeta: "Ciudad Agregada"}, nil
}

func (s *Server) UpdateNameMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Planeta: "Nombre Actualizado"}, nil
}

func (s *Server) UpdateNumberMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Planeta: "Número Actualizado "}, nil
}

func (s *Server) DeleteCityMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Mensaje que se está recibiendo: \n Planeta: %s \n Ciudad: %s \n Valor: %s", message.Planeta, message.Ciudad, message.Valor)
	return &Message{Planeta: "Ciudad Eliminada"}, nil
}