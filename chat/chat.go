package chat

import(
	"log"

	"golang.org/x/net/context"
)

type Server struct {

}

func (s *Server) SendMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Mensaje recibido: %s", message.Planeta)
	return &Message{Planeta: "Hola del servidor "}, nil
}