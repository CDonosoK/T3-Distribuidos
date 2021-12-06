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
	X int32
	Y int32
	z int32
	ultimoServidor int32
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
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatClient(conn)
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
				response, err := c.AddCityMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling SendMessage: %s", err)
				}
				//Se asumirá que nunca se agregarán ciudades duplicadas
				nuevoPlaneta := informacion{
					Planeta: message.Planeta,
					Ciudad: message.Ciudad,
					Valor: message.Valor,
					X: 0,
					Y: 0,
					z: 0,
				}
				listaPlanetas = append(listaPlanetas, nuevoPlaneta)
				log.Printf("Conectado con el servidor: %d", response.Servidor)

			// Código para actualizar el nombre de la ciudad
			case "2":
				message = UpdateName()
				response, err := c.UpdateNameMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling SendMessage: %s", err)
				}

				//Se busca el planeta y la ciudad y se actualiza el reloj de vectores
				for i := 0; i < len(listaPlanetas); i++ {
					if listaPlanetas[i].Planeta == message.Planeta && listaPlanetas[i].Ciudad == message.Ciudad {
						listaPlanetas[i].Ciudad = message.Valor
						if response.Servidor == 0 {
							listaPlanetas[i].X += 1
						}
						if response.Servidor == 1 {
							listaPlanetas[i].Y += 1
						}
						if response.Servidor == 2 {
							listaPlanetas[i].z += 1
						}
					}
				}

			// Código para actualizar el valor de la ciudad
			case "3":
				message = UpdateNumber()
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
							listaPlanetas[i].X += 1
						}
						if response.Servidor == 1 {
							listaPlanetas[i].Y += 1
						}
						if response.Servidor == 2 {
							listaPlanetas[i].z += 1
						}
					}
				}

			// Código para eliminar la ciudad
			case "4":
				message = DeleteCity()
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