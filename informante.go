package main

import(
	"log"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/CDonosoK/T3-Distribuidos/chat"

)

func AddCity(){
	var nombrePlaneta string
	fmt.Println("Ingrese el nombre del planeta: ")
	fmt.Scan(&nombrePlaneta)

	var nombreCiudad string
	fmt.Println("Ingrese el nombre de la ciudad: ")
	fmt.Scan(&nombreCiudad)

	var nuevoValor string
	fmt.Println("Ingrese el valor: ")
	fmt.Scan(&nuevoValor)
	if (nuevoValor == "") {
		nuevoValor = "0"
	}
	/*
	Conexión con los servidores fulcrum para agregar la ciudad
	*/

}

func UpdateName(){
	var nombrePlaneta string
	fmt.Println("Ingrese el nombre del planeta: ")
	fmt.Scan(&nombrePlaneta)

	var nombreCiudad string
	fmt.Println("Ingrese el nombre de la ciudad: ")
	fmt.Scan(&nombreCiudad)

	var nuevoValor string
	fmt.Println("Ingrese el nuevo nombre: ")
	fmt.Scan(&nuevoValor)

	/*
	Conexión con los servidores fulcrum para actualizar el nombre
	*/


}

func UpdateNumber(){
	var nombrePlaneta string
	fmt.Println("Ingrese el nombre del planeta: ")
	fmt.Scan(&nombrePlaneta)

	var nombreCiudad string
	fmt.Println("Ingrese el nombre de la ciudad: ")
	fmt.Scan(&nombreCiudad)

	var nuevoValor string
	fmt.Println("Ingrese el nuevo valor: ")
	fmt.Scan(&nuevoValor)
	/*
	Conexión con los servidores fulcrum para actualizar el valor
	*/


}

func DeleteCity(){
	var nombrePlaneta string
	fmt.Println("Ingrese el nombre del planeta: ")
	fmt.Scan(&nombrePlaneta)

	var nombreCiudad string
	fmt.Println("Ingrese el nombre de la ciudad: ")
	fmt.Scan(&nombreCiudad)
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

		switch decision {
		case "1":
			AddCity()
		case "2":
			UpdateName()
		case "3":
			UpdateNumber()
		case "4":
			DeleteCity()
		}

		message := chat.Message{
			Planeta: decision,
		}
		response, err := c.SendMessage(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SendMessage: %s", err)
		}
		log.Printf("Response from server: %s", response.Planeta)

		contador ++

		if (contador == 2) {
			break
		}

	}

	
}