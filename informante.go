package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/CDonosoK/T3-Distribuidos/chat"
)

type informacion struct {
	Planeta string
	Ciudad string
	Valor string

	ultimoServidor int32
}

type vectorReloj struct{
	X int32
	Y int32
	Z int32
}

var reloj = vectorReloj{
	X: 0,
	Y: 0,
	Z: 0,
}

var listaPlanetas []informacion

func AddCity() chat.Message{
	var nombrePlaneta string
	fmt.Println("Ingrese el nombre del planeta: ")
	fmt.Scan(&nombrePlaneta)

	var nombreCiudad string
	fmt.Println("Ingrese el nombre de la ciudad: ")
	fmt.Scan(&nombreCiudad)

	var nuevoValor string
	fmt.Println("Ingrese el valor: ")
	fmt.Scan(&nuevoValor)
	if (nuevoValor == " ") {
		nuevoValor = "0"
	}

	mensaje := chat.Message{
		Planeta: nombrePlaneta,
		Ciudad: nombreCiudad,
		Valor: nuevoValor,
	}

	return mensaje

}

func UpdateName() chat.Message{
	var nombrePlaneta string
	fmt.Println("Ingrese el nombre del planeta: ")
	fmt.Scan(&nombrePlaneta)

	var nombreCiudad string
	fmt.Println("Ingrese el nombre de la ciudad: ")
	fmt.Scan(&nombreCiudad)

	var nuevoValor string
	fmt.Println("Ingrese el nuevo nombre: ")
	fmt.Scan(&nuevoValor)

	mensaje := chat.Message{
		Planeta: nombrePlaneta,
		Ciudad: nombreCiudad,
		Valor: nuevoValor,
	}

	return mensaje

}

func UpdateNumber() chat.Message{
	var nombrePlaneta string
	fmt.Println("Ingrese el nombre del planeta: ")
	fmt.Scan(&nombrePlaneta)

	var nombreCiudad string
	fmt.Println("Ingrese el nombre de la ciudad: ")
	fmt.Scan(&nombreCiudad)

	var nuevoValor string
	fmt.Println("Ingrese el nuevo valor: ")
	fmt.Scan(&nuevoValor)

	mensaje := chat.Message{
		Planeta: nombrePlaneta,
		Ciudad: nombreCiudad,
		Valor: nuevoValor,
	}

	return mensaje

}

func DeleteCity() chat.Message{
	var nombrePlaneta string
	fmt.Println("Ingrese el nombre del planeta: ")
	fmt.Scan(&nombrePlaneta)

	var nombreCiudad string
	fmt.Println("Ingrese el nombre de la ciudad: ")
	fmt.Scan(&nombreCiudad)

	mensaje := chat.Message{
		Planeta: nombrePlaneta,
		Ciudad: nombreCiudad,
		Valor: "-",
	}

	return mensaje

}

func main() {

	//Conexión informantes con el servidor broker
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatClient(conn)

	//Conexión informantes con el fulcrum 1
	var conn1 *grpc.ClientConn
	conn1, err1 := grpc.Dial(":9002", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("Could not connect: %s", err1)
	}
	defer conn.Close()
	c1 := chat.NewChatClient(conn1)

	//Conexión informantes con el fulcrum 2
	var conn2 *grpc.ClientConn
	conn2, err2 := grpc.Dial(":9003", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("Could not connect: %s", err2)
	}
	defer conn.Close()
	c2 := chat.NewChatClient(conn2)

	//Conexión informantes con el fulcrum 3
	var conn3 *grpc.ClientConn
	conn3, err3 := grpc.Dial(":9004", grpc.WithInsecure())
	if err3 != nil {
		log.Fatalf("Could not connect: %s", err3)
	}
	defer conn.Close()
	c3 := chat.NewChatClient(conn3)

	salir := false

	for {
		var decision string
		fmt.Println("--- Menu Principal - Informantes ---")
		fmt.Println("¿Qué comando desea utilizar? \n [1] AddCity \n [2] UpdateName \n [3] UpdateNumber \n [4] DeleteCity \n [5] Salir")
		fmt.Scan(&decision)

		message := chat.Message{}

		switch decision {
			// Código para agregar la ciudad
			case "1":
				message = AddCity()
				message.Tipo = "0"
				response, err := c.AddCityMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling SendMessage: %s", err)
				}
				//Se asumirá que nunca se agregarán ciudades duplicadas
				nuevoPlaneta := informacion{
					Planeta: message.Planeta,
					Ciudad: message.Ciudad,
					Valor: message.Valor,

					ultimoServidor: response.Servidor,
				}
				if (response.Servidor == 0) {
					reloj.X += 1
				}
				if (response.Servidor == 1) {
					reloj.Y += 1
				}
				if (response.Servidor == 2) {
					reloj.Z += 1
				}
				listaPlanetas = append(listaPlanetas, nuevoPlaneta)
				//Conexión con los servidores Fulcrum

				//BORRAR ESA LÍNEA
				response.Servidor = 0
				if (response.Servidor == 0) {
					//Se envia el mensaje al servidor Fulcrum 1
					responsef1, errf1 := c1.AddCityF(context.Background(), &message)
					if errf1 != nil {
						log.Fatalf("Error when calling SendMessage: %s", errf1)
					}
					log.Printf("Conectado con el servidor: %d", responsef1.Servidor)

				}
				if (response.Servidor == 1) {
					//Se envia el mensaje al servidor Fulcrum 2
					responsef2, errf2 := c2.AddCityF(context.Background(), &message)
					if errf2 != nil {
						log.Fatalf("Error when calling SendMessage: %s", errf2)
					}
					log.Printf("Conectado con el servidor: %d", responsef2.Servidor)
				}
				if (response.Servidor == 2) {
					//Se envia el mensaje al servidor Fulcrum 3
					responsef3, errf3 := c3.AddCityF(context.Background(), &message)
					if errf3 != nil {
						log.Fatalf("Error when calling SendMessage: %s", errf3)
					}
					log.Printf("Conectado con el servidor: %d", responsef3.Servidor)
				}
				
				log.Printf("Conectado con el servidor: %d", response.Servidor)

			// Código para actualizar el nombre de la ciudad
			case "2":
				message = UpdateName()
				message.Tipo = "1"
				response, err := c.UpdateNameMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling SendMessage: %s", err)
				}

				//Se busca el planeta y la ciudad y se actualiza el reloj de vectores
				for i := 0; i < len(listaPlanetas); i++ {
					if listaPlanetas[i].Planeta == message.Planeta && listaPlanetas[i].Ciudad == message.Ciudad {
						listaPlanetas[i].Ciudad = message.Valor
						if response.Servidor == 0 {
							reloj.X += 1
						}
						if response.Servidor == 1 {
							reloj.Y += 1
						}
						if response.Servidor == 2 {
							reloj.Z += 1
						}
					}
				}

				log.Printf("Conectado con el servidor: %d", response.Servidor)

			// Código para actualizar el valor de la ciudad
			case "3":
				message = UpdateNumber()
				message.Tipo = "2"
				response, err := c.UpdateNumberMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling SendMessage: %s", err)
				}
				log.Printf("Response from server: %d", response.Servidor)
				//Se busca el planeta y la ciudad y se actualiza el reloj de vectores
				for i := 0; i < len(listaPlanetas); i++ {
					if listaPlanetas[i].Planeta == message.Planeta && listaPlanetas[i].Ciudad == message.Ciudad {
						listaPlanetas[i].Valor = message.Valor
						if response.Servidor == 0 {
							reloj.X += 1
						}
						if response.Servidor == 1 {
							reloj.Y += 1
						}
						if response.Servidor == 2 {
							reloj.Z += 1
						}
					}
				}
				//Conexión con los servidores Fulcrum
				log.Printf("Conectado con el servidor: %d", response.Servidor)

			// Código para eliminar la ciudad
			case "4":
				message = DeleteCity()
				message.Tipo = "3"
				response, err := c.DeleteCityMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling SendMessage: %s", err)
				}
				log.Printf("Response from server: %d", response.Servidor)
				//Se busca el planeta y la ciudad y se guarda la posición dónde se encuentra
				var i = 0
				for i = 0; i < len(listaPlanetas); i++ {
					if listaPlanetas[i].Planeta == message.Planeta && listaPlanetas[i].Ciudad == message.Ciudad {
						break
					}
				}
				//Se elimina el registro de la memoria
				listaPlanetas[i] = listaPlanetas[len(listaPlanetas)-1]
				listaPlanetas = listaPlanetas[:len(listaPlanetas)-1]

				//Conexión con los servidores Fulcrum
				log.Printf("Conectado con el servidor: %d", response.Servidor)

			// Código para salir del programa
			case "5":
				salir = true
			}
			
			//Si se quiere ver la información guardada en memoria, descomentar la siguiente línea
			fmt.Println(listaPlanetas)

			if salir {
				break
			}

	}

	
}