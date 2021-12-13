# T3-Distribuidos
## Star-Wars

## INTEGRANTES:
- Nazareth Díaz - 201873560-4
- Jael Toledo - 201873543-4
- Clemente Donoso - 201873546-9

Las máquinas virtuales están asociadas de la siguiente manera:

|Máquina | Proceso |
| ----- | ----- |
| dist69 | brokerMosEisley |
| dist70 | servidorFulcrum 1, informante 1 |
| dist71 | servidorFulcrum 2, informante 2 |
| dist72 | servidorFulcrum 3, leia |

Es de suma importancia que al ejecutar las máquinas se deban realizar en el siguiente órden

|Máquina | Lugar |
| ----- | ----- |
| dist69 | 1 |
| dist70 | 2 |
| dist71 | 3 |
| dist72 | 4 |

En cada máquina aparecerán las instrucciones sobre como continuar, para que así el proceso sea lo más legible y entendible posible.

### INSTRUCCIONES:
Es importante aclarar que las máquinas dis70, dis71 y dis72 deben ejecutar los comandos en distintas terminales.
- Se debe ejecutar en la máquina dist69 el comando: 
``` go run brokerMosEisley.go```
- Se debe ejecutar en la máquina dist70 el comando: 
``` go run sercidorFulcrum1.go``` 
``` go run informante.go ``` 
- Se debe ejecutar en la máquina dist71 el comando: 
``` go run sercidorFulcrum2.go ```  
``` go run informante.go``` 
- Se debe ejecutar en la máquina dist72 el comando: 
``` go run sercidorFulcrum3.go```  
``` go run leia.go``` 

## CONSIDERACIONES
- Se asume que tanto los informantes, como leia son inteligentes y seguirán las instrucciones que aparecen en pantalla.
- Se asumirá que los informantes al agregar una ciudad no agregarán ciudades respetidas.
- Al momento de actualizar un valor (ya sea ciudad o valor), se asume que se indicará un lugar existente.
- El reloj está pensado como un reloj universal, donde cada componente corresponde al Servidor Fulcrum al cual se le están realizando modificaciones.
- La consistencia read your writes se hace en base a lo que envía el informante y en base a lo que recibe de vuelta, donde deben ser los mismos datos.
- El brokerMosEisley envía de respuesta un {0,1,2} si se conecta aleatoriamente con el servidor Fulcrum 1, 2 o 3 respectivamente, y luego el informante realiza la conexión pertinente.

RECORDAR BORRAR LO SIGUIENTE:

- Para actualizar el archivo proto se debe ejecutar el siguiente comando: protoc --go_out=plugins=grpc:chat chat.proto
- Ejecutar el broker con el comando: go run brokerMosEisley.go
- Ejecutar los informantes con el comando: go run informante.go