package chat

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct {

}

type infoPlaneta struct {
	NomPlaneta string
	ciudad 	string
	valor string
}

type relojPlaneta struct {
	Planeta string
	X int32
	Y int32
	Z int32
}

var listaPlaneta []infoPlaneta
var listaReloj []relojPlaneta

func (s *Server) AddCityF(ctx context.Context, message *Message) (*Message, error) {
	// SE OBTIENE LA INFORMACIÓN
	planeta := message.Planeta

	nuevoPlaneta := infoPlaneta{
		NomPlaneta: message.Planeta,
		ciudad: message.Ciudad,
		valor: message.Valor,
	}

	// SE ACTUALIZA EL RELOJ
	var planetaEsta = false
	if (len(listaReloj) == 0) {
		listaReloj = append(listaReloj, relojPlaneta{Planeta: planeta, X: 0, Y: 0, Z: 0})
		if message.Servidor == 0 {
			listaReloj[0].X += 1
		}
		if message.Servidor == 1 {
			listaReloj[0].Y += 1
		}
		if message.Servidor == 2 {
			listaReloj[0].Z += 1
		}
	} else {
		for i := 0; i < len(listaReloj); i++ {
			if listaReloj[i].Planeta == planeta {
				planetaEsta = true
				if message.Servidor == 0 {
					listaReloj[i].X += 1
				}
				if message.Servidor == 1 {
					listaReloj[i].Y += 1
				}
				if message.Servidor == 2 {
					listaReloj[i].Z += 1
				}
			}
		}
		if !planetaEsta {
			listaReloj = append(listaReloj, relojPlaneta{Planeta: planeta, X: 0, Y: 0, Z: 0})
			if message.Servidor == 0 {
				listaReloj[len(listaReloj)-1].X += 1
			}
			if message.Servidor == 1 {
				listaReloj[len(listaReloj)-1].Y += 1
			}
			if message.Servidor == 2 {
				listaReloj[len(listaReloj)-1].Z += 1
			}
		}
	}
	

	listaPlaneta = append(listaPlaneta, nuevoPlaneta)
	fmt.Println(listaReloj)

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
	log.Println("\n------------------------------------------------------")
	log.Println(logMessage)
	log.Println("\n------------------------------------------------------")

	// Actualizar el return

	return &Message{Planeta: "RECIBIDO"}, nil
}

func (s *Server) AddCityMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))	
	logMessage := "AddCity " + message.Planeta + " " + message.Ciudad + " " + message.Valor + "\n"

	log.Println("\n------------------------------------------------------")
	log.Println(logMessage)
	log.Println("\n------------------------------------------------------")
	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Valor: message.Valor, Servidor: serverElegido}, nil
}

func (s *Server) UpdateNameF(ctx context.Context, message *Message) (*Message, error) {
	planeta := message.Planeta
	ciudad := message.Ciudad
	valor := message.Valor

	directorio1 := "./Logs/" + planeta + ".txt"
	directorio2 := "./Registros Planetarios/" + planeta + ".txt"

	for i := 0; i < len(listaPlaneta); i++ {
		if listaPlaneta[i].NomPlaneta == planeta && listaPlaneta[i].ciudad == ciudad {
			listaPlaneta[i].ciudad = valor
		}
	}
	os.Remove(directorio2)

	// SE ESCRIBE EN EL LOGS - IGUAL PARA TODAS LAS FUNCIONES
	logMessage := "UpdateName " + message.Planeta + " " + message.Ciudad + " " + message.Valor + "\n"
	f, err := os.OpenFile(directorio1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(logMessage); err != nil {
		log.Println(err)
	}

	// SE REESCRIBE EL REGISTRO
	for i := 0; i < len(listaPlaneta); i++ {
		p := listaPlaneta[i].NomPlaneta
		c := listaPlaneta[i].ciudad
		v := listaPlaneta[i].valor
		registroMessage := p + " " + c + " " + v + "\n"
		f1, err1 := os.OpenFile(directorio2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			log.Println(err1)
		}
		defer f1.Close()
		if _, err := f1.WriteString(registroMessage); err != nil {
			log.Println(err)
		}
	}

	// SE ACTUALIZA EL RELOJ
	var planetaEsta = false
	if (len(listaReloj) == 0) {
		listaReloj = append(listaReloj, relojPlaneta{Planeta: planeta, X: 0, Y: 0, Z: 0})
		if message.Servidor == 0 {
			listaReloj[0].X += 1
		}
		if message.Servidor == 1 {
			listaReloj[0].Y += 1
		}
		if message.Servidor == 2 {
			listaReloj[0].Z += 1
		}
	} else {
		for i := 0; i < len(listaReloj); i++ {
			if listaReloj[i].Planeta == planeta {
				planetaEsta = true
				if message.Servidor == 0 {
					listaReloj[i].X += 1
				}
				if message.Servidor == 1 {
					listaReloj[i].Y += 1
				}
				if message.Servidor == 2 {
					listaReloj[i].Z += 1
				}
			}
		}
		if !planetaEsta {
			listaReloj = append(listaReloj, relojPlaneta{Planeta: planeta, X: 0, Y: 0, Z: 0})
			if message.Servidor == 0 {
				listaReloj[len(listaReloj)-1].X += 1
			}
			if message.Servidor == 1 {
				listaReloj[len(listaReloj)-1].Y += 1
			}
			if message.Servidor == 2 {
				listaReloj[len(listaReloj)-1].Z += 1
			}
		}
	}
	log.Println("\n------------------------------------------------------")
	log.Println(logMessage)
	log.Println("\n------------------------------------------------------")


	serverElegido := int32(rand.Intn(3))

	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Valor: message.Valor, Servidor: serverElegido}, nil
}

func (s *Server) UpdateNameMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))

	logMessage := "UpdateName " + message.Planeta + " " + message.Ciudad + " " + message.Valor + "\n"
	log.Println("\n------------------------------------------------------")
	log.Println(logMessage)
	log.Println("\n------------------------------------------------------")

	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Valor: message.Valor, Servidor: serverElegido}, nil
}

func (s *Server) UpdateNumberF(ctx context.Context, message *Message) (*Message, error) {
	planeta := message.Planeta
	ciudad := message.Ciudad
	valor := message.Valor

	directorio1 := "./Logs/" + planeta + ".txt"
	directorio2 := "./Registros Planetarios/" + planeta + ".txt"

	for i := 0; i < len(listaPlaneta); i++ {
		if listaPlaneta[i].NomPlaneta == planeta && listaPlaneta[i].ciudad == ciudad {
			listaPlaneta[i].valor = valor
		}
	}
	os.Remove(directorio2)

	// SE ESCRIBE EN EL LOGS - IGUAL PARA TODAS LAS FUNCIONES
	logMessage := "UpdateNumber " + message.Planeta + " " + message.Ciudad + " " + message.Valor + "\n"
	f, err := os.OpenFile(directorio1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(logMessage); err != nil {
		log.Println(err)
	}

	// SE REESCRIBE EL REGISTRO
	for i := 0; i < len(listaPlaneta); i++ {
		p := listaPlaneta[i].NomPlaneta
		c := listaPlaneta[i].ciudad
		v := listaPlaneta[i].valor
		registroMessage := p + " " + c + " " + v + "\n"
		f1, err1 := os.OpenFile(directorio2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			log.Println(err1)
		}
		defer f1.Close()
		if _, err := f1.WriteString(registroMessage); err != nil {
			log.Println(err)
		}
	}

	// SE ACTUALIZA EL RELOJ
	var planetaEsta = false
	if (len(listaReloj) == 0) {
		listaReloj = append(listaReloj, relojPlaneta{Planeta: planeta, X: 0, Y: 0, Z: 0})
		if message.Servidor == 0 {
			listaReloj[0].X += 1
		}
		if message.Servidor == 1 {
			listaReloj[0].Y += 1
		}
		if message.Servidor == 2 {
			listaReloj[0].Z += 1
		}
	} else {
		for i := 0; i < len(listaReloj); i++ {
			if listaReloj[i].Planeta == planeta {
				planetaEsta = true
				if message.Servidor == 0 {
					listaReloj[i].X += 1
				}
				if message.Servidor == 1 {
					listaReloj[i].Y += 1
				}
				if message.Servidor == 2 {
					listaReloj[i].Z += 1
				}
			}
		}
		if !planetaEsta {
			listaReloj = append(listaReloj, relojPlaneta{Planeta: planeta, X: 0, Y: 0, Z: 0})
			if message.Servidor == 0 {
				listaReloj[len(listaReloj)-1].X += 1
			}
			if message.Servidor == 1 {
				listaReloj[len(listaReloj)-1].Y += 1
			}
			if message.Servidor == 2 {
				listaReloj[len(listaReloj)-1].Z += 1
			}
		}
	}
	
	log.Println("\n------------------------------------------------------")
	log.Println(logMessage)
	log.Println("\n------------------------------------------------------")


	serverElegido := int32(rand.Intn(3))

	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Valor: message.Valor, Servidor: serverElegido}, nil
}

func (s *Server) UpdateNumberMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))

	logMessage := "UpdateNumber " + message.Planeta + " " + message.Ciudad + " " + message.Valor + "\n"
	log.Println("\n------------------------------------------------------")
	log.Println(logMessage)
	log.Println("\n------------------------------------------------------")

	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Valor: message.Valor, Servidor: serverElegido}, nil
}

func (s *Server) DeleteCityF(ctx context.Context, message *Message) (*Message, error) {
	planeta := message.Planeta
	ciudad := message.Ciudad

	directorio1 := "./Logs/" + planeta + ".txt"
	directorio2 := "./Registros Planetarios/" + planeta + ".txt"

	var i = 0;
	for i = 0; i < len(listaPlaneta); i++ {
		if listaPlaneta[i].NomPlaneta == planeta && listaPlaneta[i].ciudad == ciudad {
			break
		}
	}
	listaPlaneta[i] = listaPlaneta[len(listaPlaneta)-1]
	listaPlaneta = listaPlaneta[:len(listaPlaneta)-1]
	os.Remove(directorio2)

	// SE ESCRIBE EN EL LOGS - IGUAL PARA TODAS LAS FUNCIONES
	logMessage := "DeleteCity " + message.Planeta + " " + message.Ciudad + "\n"
	f, err := os.OpenFile(directorio1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(logMessage); err != nil {
		log.Println(err)
	}

	// SE REESCRIBE EL REGISTRO
	for i := 0; i < len(listaPlaneta); i++ {
		p := listaPlaneta[i].NomPlaneta
		c := listaPlaneta[i].ciudad
		v := listaPlaneta[i].valor
		registroMessage := p + " " + c + " " + v + "\n"
		f1, err1 := os.OpenFile(directorio2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			log.Println(err1)
		}
		defer f1.Close()
		if _, err := f1.WriteString(registroMessage); err != nil {
			log.Println(err)
		}
	}

	// SE ACTUALIZA EL RELOJ
	var planetaEsta = false
	if (len(listaReloj) == 0) {
		listaReloj = append(listaReloj, relojPlaneta{Planeta: planeta, X: 0, Y: 0, Z: 0})
		if message.Servidor == 0 {
			listaReloj[0].X += 1
		}
		if message.Servidor == 1 {
			listaReloj[0].Y += 1
		}
		if message.Servidor == 2 {
			listaReloj[0].Z += 1
		}
	} else {
		for i := 0; i < len(listaReloj); i++ {
			if listaReloj[i].Planeta == planeta {
				planetaEsta = true
				if message.Servidor == 0 {
					listaReloj[i].X += 1
				}
				if message.Servidor == 1 {
					listaReloj[i].Y += 1
				}
				if message.Servidor == 2 {
					listaReloj[i].Z += 1
				}
			}
		}
		if !planetaEsta {
			listaReloj = append(listaReloj, relojPlaneta{Planeta: planeta, X: 0, Y: 0, Z: 0})
			if message.Servidor == 0 {
				listaReloj[len(listaReloj)-1].X += 1
			}
			if message.Servidor == 1 {
				listaReloj[len(listaReloj)-1].Y += 1
			}
			if message.Servidor == 2 {
				listaReloj[len(listaReloj)-1].Z += 1
			}
		}
	}
	
	log.Println("\n------------------------------------------------------")
	log.Println(logMessage)
	log.Println("\n------------------------------------------------------")


	serverElegido := int32(rand.Intn(3))

	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Valor: message.Valor, Servidor: serverElegido}, nil
}

func (s *Server) DeleteCityMessage(ctx context.Context, message *Message) (*Message, error) {
	serverElegido := int32(rand.Intn(3))

	logMessage := "DeleteCity " + message.Planeta + " " + message.Ciudad + "\n"
	log.Println("\n------------------------------------------------------")
	log.Println(logMessage)
	log.Println("\n------------------------------------------------------")

	return &Message{Planeta: message.Planeta, Ciudad: message.Ciudad, Servidor: serverElegido}, nil
}

func (s *Server) ObtenerNumeroRebeldesFulcrum(ctx context.Context, message *DeLeia) (*ParaLeia, error) {
	log.Printf("")
	log.Println("\n------------------------------------------------------")
	log.Printf("Mensaje que se está recibiendo desde Broker: \n ~~~ Leia quiere saber los rebeldes de \n Planeta: "+message.Planeta +"  Ciudad: "+ message.Ciudad)
	log.Printf("Buscando cantidad de rebeldes...")
	numRebeldes := int32(0)
	for i := 0; i < len(listaPlaneta); i++ {
		if message.Planeta == listaPlaneta[i].NomPlaneta{
			nRebeldesS,_ := strconv.Atoi(listaPlaneta[i].valor)
			numRebeldes = int32(nRebeldesS)
		}
	}

	x := int32(-1)
	y := int32(-1)
	z := int32(-1)
	for i := 0; i <len(listaReloj); i++{
		if message.Planeta == listaReloj[i].Planeta{
			x = listaReloj[i].X
			y = listaReloj[i].Y 
			z = listaReloj[i].Z
		}
	}
	log.Println("\n------------------------------------------------------")
	return &ParaLeia{CantRebeldes: numRebeldes, X: x, Y: y, Z: z, Servidor: -1}, nil
}


func (s *Server) Merge(ctx context.Context, message *Merge) (*Merge, error) {


	return &Merge{}, nil
}


func (s *Server) ObtenerNumeroRebeldesBroker(ctx context.Context, message *DeLeia) (*ParaLeia, error) {
	serverElegido := int32(rand.Intn(3))
	serverNombre := " "

	//Por defecto se va con el 1
	serverNombre = "Servidor Fulcrum 1"
	var conn1 *grpc.ClientConn
	conn1, err1 := grpc.Dial("dist70:9002", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("Could not connect: %s", err1)
	}
	defer conn1.Close()
	c := NewChatClient(conn1)

	if (serverElegido == 1){
		serverNombre = "Servidor Fulcrum 2"
		var conn2 *grpc.ClientConn
		conn2, err2 := grpc.Dial("dist71:9003", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("Could not connect: %s", err2)
		}
		defer conn2.Close()
		c = NewChatClient(conn2)
	}
	if (serverElegido == 2){
		serverNombre = "Servidor Fulcrum 3"
		var conn3 *grpc.ClientConn
		conn3, err3 := grpc.Dial("dist72:9004", grpc.WithInsecure())
		if err3 != nil {
			log.Fatalf("Could not connect: %s", err3)
		}
		defer conn3.Close()
		c = NewChatClient(conn3)
	}
	log.Println("\n------------------------------------------------------")
	log.Printf("~~Leia solicita la cantidad de rebeldes en %s, %s \n", message.Planeta, message.Ciudad)
	log.Printf("~~El servidor escogido aleatoriamente es: %s", serverNombre)
	log.Println("\n------------------------------------------------------")

	respuestaFulcrum, err := c.ObtenerNumeroRebeldesFulcrum(context.Background(), message)
	
	if err != nil {
		log.Fatalf("Fulcrum no ha mandado info. %s",err)
	}

	return &ParaLeia{CantRebeldes: respuestaFulcrum.CantRebeldes, X: respuestaFulcrum.X, Y: respuestaFulcrum.Y, Z: respuestaFulcrum.Z, Servidor: serverElegido}, nil

}

