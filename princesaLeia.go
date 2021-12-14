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
	Valor int32
	X int32
	Y int32
	Z int32
	UltimoServidor int32
}

var ArrayInfo []informacion
var planeta string
var ciudad string



func main(){
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
		fmt.Println("--- Menu Principal - Leia ---")
		fmt.Println("¿Qué comando desea utilizar? \n [1] Obtener el número de rebeldes \n [2] Salir")
		fmt.Scanln(&decision)		

		switch decision {
			case "1":
				fmt.Println("Ingrese el nombre del planeta: ")
				fmt.Scanln(&planeta)

				fmt.Println("Ingrese la ciudad a buscar")
				fmt.Scanln(&ciudad)
				
				x := int32(-1)
				y := int32(-1)
				z := int32(-1)
				fmt.Println("Llega aca")
				for i := 0; i < len(ArrayInfo); i++ {
					if planeta == ArrayInfo[i].Planeta && ciudad == ArrayInfo[i].Ciudad{
						x = ArrayInfo[i].X
						y = ArrayInfo[i].Y
						z = ArrayInfo[i].Z
					}
				}

				fmt.Println("Llega aca2")
				mensajeDeLeia := chat.DeLeia{
					X: x, 
					Y: y, 
					Z: z, 
					Planeta: planeta, 
					Ciudad: ciudad,
				}

				fmt.Println("Llega aca3")
				respuestaALeia, err := c.ObtenerNumeroRebeldesBroker(context.Background(), &mensajeDeLeia)
				if err != nil {
					log.Fatalf("Error when calling ObtenerNumeroRebeldesBroker: %s", err)
				}			
				flag := 0
				
				fmt.Println("Llega aca4")

				for i := 0; i < len(ArrayInfo); i++ {
					if planeta == ArrayInfo[i].Planeta && ciudad == ArrayInfo[i].Ciudad{
						flag = 1
						if respuestaALeia.X >= ArrayInfo[i].X && respuestaALeia.Y >= ArrayInfo[i].Y && respuestaALeia.Z >= ArrayInfo[i].Z{
							//Se actualiza porque la respuesta esta mas actualizada

							ArrayInfo[i].Valor = respuestaALeia.CantRebeldes
							ArrayInfo[i].X = respuestaALeia.X
							ArrayInfo[i].Y = respuestaALeia.Y 
							ArrayInfo[i].Z = respuestaALeia.Z
							ArrayInfo[i].UltimoServidor = respuestaALeia.Servidor

							fmt.Println("La cantidad de rebeldes en")
							fmt.Println(ciudad, planeta)
							fmt.Println( respuestaALeia.CantRebeldes)
							fmt.Println("\n")
							fmt.Println("El reloj del planeta es ")
							fmt.Println(ArrayInfo[i].X, ArrayInfo[i].Y, ArrayInfo[i].Z)
						}
					}
				}

				if flag == 0{
					nuevoStruct := informacion{
						Planeta: planeta,
						Ciudad: ciudad,
						Valor: respuestaALeia.CantRebeldes,
						X: respuestaALeia.X,
						Y: respuestaALeia.Y,
						Z: respuestaALeia.Z,
						UltimoServidor: respuestaALeia.Servidor,
					}
					ArrayInfo = append(ArrayInfo, nuevoStruct)

					fmt.Println("La cantidad de rebeldes en")
					fmt.Println(ciudad, planeta)					
					fmt.Println(nuevoStruct.Valor)
					fmt.Println("\n")
					fmt.Println("El reloj del planeta es ")
					fmt.Println(nuevoStruct.X, nuevoStruct.Y, nuevoStruct.Z)
				}

				

			case "2":
				salir = true
		}
		if salir {
			break
		}
	}
}