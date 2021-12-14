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
	Planeta string
	X int32
	Y int32
	Z int32
}

var listaReloj []vectorReloj
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
	defer conn2.Close()
	c2 := chat.NewChatClient(conn2)

	//Conexión informantes con el fulcrum 3
	var conn3 *grpc.ClientConn
	conn3, err3 := grpc.Dial(":9004", grpc.WithInsecure())
	if err3 != nil {
		log.Fatalf("Could not connect: %s", err3)
	}
	defer conn3.Close()
	c3 := chat.NewChatClient(conn3)

	salir := false

	for {
		var decision string
		fmt.Println("--- Menu Principal - Informantes ---")
		fmt.Println("¿Qué comando desea utilizar? \n [1] AddCity \n [2] UpdateName \n [3] UpdateNumber \n [4] DeleteCity \n [5] Salir")
		fmt.Scan(&decision)

		message := chat.Message{}

		switch decision {
			// CÓDIGO PARA AGREGAR CIUDAD
			case "1":
				message = AddCity()
				message.Tipo = "0"
				response, err := c.AddCityMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling AddCityMessage: %s", err)
				}
				//CONSISTENCIA READ YOUR WRITES
				if (message.Planeta == response.Planeta && message.Ciudad == response.Ciudad && message.Valor == response.Valor) {
					log.Println("Existe consistencia Read Your Writes")
				}

				// SE ASUME QUE NO SE AGREGARÁN CIUDADES DUPLICADAS
				nuevoPlaneta := informacion{
					Planeta: message.Planeta,
					Ciudad: message.Ciudad,
					Valor: message.Valor,

					ultimoServidor: response.Servidor,
				}

				// SE ACTUALIZA EL RELOJ
				var planetaEsta = false
				if (len(listaReloj) == 0) {
					listaReloj = append(listaReloj, vectorReloj{Planeta: message.Planeta, X: 0, Y: 0, Z: 0})
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
						if listaReloj[i].Planeta == message.Planeta {
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
						listaReloj = append(listaReloj, vectorReloj{Planeta: message.Planeta, X: 0, Y: 0, Z: 0})
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

				listaPlanetas = append(listaPlanetas, nuevoPlaneta)
				//CONEXIÓN CON LOS SERVIDORES FULCRUM

				if (response.Servidor == 0) {
					// SE ENVÍA EL MENSAJE A FULCRUM 1
					responsef1, errf1 := c1.AddCityF(context.Background(), &message)
					if errf1 != nil {
						log.Fatalf("Error when calling AddCityF: %s", errf1)
					}
					log.Printf("Conectado con el servidor: %d", responsef1.Servidor)

				}
				if (response.Servidor == 1) {
					// SE ENVÍA EL MENSAJE A FULCRUM 2
					responsef2, errf2 := c2.AddCityF(context.Background(), &message)
					if errf2 != nil {
						log.Fatalf("Error when calling AddCityF: %s", errf2)
					}
					log.Printf("Conectado con el servidor: %d", responsef2.Servidor)
				}
				if (response.Servidor == 2) {
					// SE ENVÍA EL MENSAJE A FULCRUM 3
					responsef3, errf3 := c3.AddCityF(context.Background(), &message)
					if errf3 != nil {
						log.Fatalf("Error when calling AddCityF: %s", errf3)
					}
					log.Printf("Conectado con el servidor: %d", responsef3.Servidor)
				}

			// CÓDIGO PARA ACTUALIZAR NOMBRE DE LA CIUDAD
			case "2":
				message = UpdateName()
				message.Tipo = "1"
				response, err := c.UpdateNameMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling UpdateNameMessage: %s", err)
				}

				//CONSISTENCIA READ YOUR WRITES
				if (message.Planeta == response.Planeta && message.Ciudad == response.Ciudad && message.Valor == response.Valor) {
					log.Println("Existe consistencia Read Your Writes")
				}

				// SE BUSCA EL PLANETA Y LA CIUDAD PARA ACTUALIZAR ESE DATO
				for i := 0; i < len(listaPlanetas); i++ {
					if listaPlanetas[i].Planeta == message.Planeta && listaPlanetas[i].Ciudad == message.Ciudad {
						listaPlanetas[i].Ciudad = message.Valor
						// SE ACTUALIZA EL RELOJ
						var planetaEsta = false
						if (len(listaReloj) == 0) {
							listaReloj = append(listaReloj, vectorReloj{Planeta: message.Planeta, X: 0, Y: 0, Z: 0})
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
								if listaReloj[i].Planeta == message.Planeta {
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
								listaReloj = append(listaReloj, vectorReloj{Planeta: message.Planeta, X: 0, Y: 0, Z: 0})
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
					}
				}
				// CONEXIÓN CON LOS SERVIDORES FULCRUM

				if (response.Servidor == 0) {
					// SE ENVÍA EL MENSAJE A FULCRUM 1
					responsef1, errf1 := c1.UpdateNameF(context.Background(), &message)
					if errf1 != nil {
						log.Fatalf("Error when calling UpdateNameF: %s", errf1)
					}
					log.Printf("Conectado con el servidor: %d", responsef1.Servidor)

				}
				if (response.Servidor == 1) {
					// SE ENVÍA EL MENSAJE A FULCRUM 2
					responsef2, errf2 := c2.UpdateNameF(context.Background(), &message)
					if errf2 != nil {
						log.Fatalf("Error when calling UpdateNameF: %s", errf2)
					}
					log.Printf("Conectado con el servidor: %d", responsef2.Servidor)
				}
				if (response.Servidor == 2) {
					// SE ENVÍA EL MENSAJE A FULCRUM 3
					responsef3, errf3 := c3.UpdateNameF(context.Background(), &message)
					if errf3 != nil {
						log.Fatalf("Error when calling UpdateNameF: %s", errf3)
					}
					log.Printf("Conectado con el servidor: %d", responsef3.Servidor)
				}

			// CÓDIGO PARA ACTUALIZAR NUMERO DE REBELDES
			case "3":
				message = UpdateNumber()
				message.Tipo = "2"
				response, err := c.UpdateNumberMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling UpdateNumberMessage: %s", err)
				}
				log.Printf("Response from server: %d", response.Servidor)

				//CONSISTENCIA READ YOUR WRITES
				if (message.Planeta == response.Planeta && message.Ciudad == response.Ciudad && message.Valor == response.Valor) {
					log.Println("Existe consistencia Read Your Writes")
				}
				
				// SE BUSCA EL PLANETA Y LA CIUDAD PARA ACTUALIZAR ESE DATO
				for i := 0; i < len(listaPlanetas); i++ {
					if listaPlanetas[i].Planeta == message.Planeta && listaPlanetas[i].Ciudad == message.Ciudad {
						listaPlanetas[i].Valor = message.Valor
						// SE ACTUALIZA EL RELOJ
						var planetaEsta = false
						if (len(listaReloj) == 0) {
							listaReloj = append(listaReloj, vectorReloj{Planeta: message.Planeta, X: 0, Y: 0, Z: 0})
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
								if listaReloj[i].Planeta == message.Planeta {
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
								listaReloj = append(listaReloj, vectorReloj{Planeta: message.Planeta, X: 0, Y: 0, Z: 0})
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
					}
				}
				// CONEXIÓN CON LOS SERVIDORES FULCRUM

				if (response.Servidor == 0) {
					// SE ENVÍA EL MENSAJE A FULCRUM 1
					responsef1, errf1 := c1.UpdateNumberF(context.Background(), &message)
					if errf1 != nil {
						log.Fatalf("Error when calling UpdateNumberF: %s", errf1)
					}
					log.Printf("Conectado con el servidor: %d", responsef1.Servidor)

				}
				if (response.Servidor == 1) {
					// SE ENVÍA EL MENSAJE A FULCRUM 2
					responsef2, errf2 := c2.UpdateNumberF(context.Background(), &message)
					if errf2 != nil {
						log.Fatalf("Error when calling UpdateNumberF: %s", errf2)
					}
					log.Printf("Conectado con el servidor: %d", responsef2.Servidor)
				}
				if (response.Servidor == 2) {
					// SE ENVÍA EL MENSAJE A FULCRUM 3
					responsef3, errf3 := c3.UpdateNumberF(context.Background(), &message)
					if errf3 != nil {
						log.Fatalf("Error when calling UpdateNumberF: %s", errf3)
					}
					log.Printf("Conectado con el servidor: %d", responsef3.Servidor)
				}


			// CÓDIGO PARA ELIMINAR LA CIUDAD
			case "4":
				message = DeleteCity()
				message.Tipo = "3"
				response, err := c.DeleteCityMessage(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling DeleteCityMessage: %s", err)
				}
				log.Printf("Response from server: %d", response.Servidor)

				//CONSISTENCIA READ YOUR WRITES
				if (message.Planeta == response.Planeta && message.Ciudad == response.Ciudad) {
					log.Println("Existe consistencia Read Your Writes")
				}

				// SE BUSCA EL PLANETA Y LA CIUDAD PARA ELIMINAR ESE DATO
				var i = 0
				for i = 0; i < len(listaPlanetas); i++ {
					if listaPlanetas[i].Planeta == message.Planeta && listaPlanetas[i].Ciudad == message.Ciudad {
						// SE ACTUALIZA EL RELOJ
						var planetaEsta = false
						if (len(listaReloj) == 0) {
							listaReloj = append(listaReloj, vectorReloj{Planeta: message.Planeta, X: 0, Y: 0, Z: 0})
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
								if listaReloj[i].Planeta == message.Planeta {
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
								listaReloj = append(listaReloj, vectorReloj{Planeta: message.Planeta, X: 0, Y: 0, Z: 0})
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
					}
				}

				// SE ELIMINA EL REGISTRO DE LA MEMORIA
				listaPlanetas[i] = listaPlanetas[len(listaPlanetas)-1]
				listaPlanetas = listaPlanetas[:len(listaPlanetas)-1]

				// CONEXIÓN CON LOS SERVIDORES FULCRUM

				if (response.Servidor == 0) {
					// SE ENVÍA EL MENSAJE A FULCRUM 1
					responsef1, errf1 := c1.DeleteCityF(context.Background(), &message)
					if errf1 != nil {
						log.Fatalf("Error when calling DeleteCityF: %s", errf1)
					}
					log.Printf("Conectado con el servidor: %d", responsef1.Servidor)

				}
				if (response.Servidor == 1) {
					// SE ENVÍA EL MENSAJE A FULCRUM 2
					responsef2, errf2 := c2.DeleteCityF(context.Background(), &message)
					if errf2 != nil {
						log.Fatalf("Error when calling DeleteCityF: %s", errf2)
					}
					log.Printf("Conectado con el servidor: %d", responsef2.Servidor)
				}
				if (response.Servidor == 2) {
					// SE ENVÍA EL MENSAJE A FULCRUM 3
					responsef3, errf3 := c3.DeleteCityF(context.Background(), &message)
					if errf3 != nil {
						log.Fatalf("Error when calling DeleteCityF: %s", errf3)
					}
					log.Printf("Conectado con el servidor: %d", responsef3.Servidor)
				}

			// SE SALE DE LA APLICACIÓN
			case "5":
				salir = true
			}
			
			//Si se quiere ver la información guardada en memoria, descomentar la siguiente línea
			//mt.Println(listaPlanetas)

			if salir {
				break
			}

	}

	
}