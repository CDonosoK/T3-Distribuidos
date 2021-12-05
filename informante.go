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
	ultimoServidor string
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
	/*
	Conexión con los servidores fulcrum para agregar la ciudad
	*/

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

	/*
	Conexión con los servidores fulcrum para actualizar el nombre
	*/


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
	/*
	Conexión con los servidores fulcrum para actualizar el valor
	*/


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
	/*
	Conexión con los servidores fulcrum para eliminar la ciudad
	*/


}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatClient(conn)
	contador := 0

	for {
		var decision string
		fmt.Println("--- Menu Principal - Informantes ---")
		fmt.Println("¿Qué comando desea utilizar? \n [1] AddCity \n [2] UpdateName \n [3] UpdateNumber \n [4] DeleteCity")
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
				log.Printf("Response from server: %d", response.Server)

			// Código para actualizar el nombre de la ciudad
			case "2":
				message = UpdateName()
				response, err := c.UpdateNameMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling SendMessage: %s", err)
				}
				log.Printf("Response from server: %d", response.Server)

			// Código para actualizar el valor de la ciudad
			case "3":
				message = UpdateNumber()
				response, err := c.UpdateNumberMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling SendMessage: %s", err)
				}
				log.Printf("Response from server: %d", response.Server)

			// Código para eliminar la ciudad
			case "4":
				message = DeleteCity()
				response, err := c.DeleteCityMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling SendMessage: %s", err)
				}
				log.Printf("Response from server: %d", response.Server)
			}

		contador ++

		if (contador == 4) {
			break
		}

	}

	
}