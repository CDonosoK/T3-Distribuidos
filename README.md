# T3-Distribuidos
## Star-Wars

## INTEGRANTES:
- Nazareth Díaz - 201873560-4
- Jael Toledo - 201873543-4
- Clemente Donoso - 201873546-9

Las máquinas virtuales están asociadas de la siguiente manera:

|Máquina | Proceso |
| ----- | ----- |
| dist69 | X |
| dist70 | X |
| dist71 | X |
| dist72 | X |

Es de suma importancia que al ejecutar las máquinas se deban realizar en el siguiente órden

|Máquina | Lugar |
| ----- | ----- |
| dist69 | X |
| dist70 | X |
| dist71 | X |
| dist72 | X |

En cada máquina aparecerán las instrucciones sobre como continuar, para que así el proceso sea lo más legible y entendible posible.

### INSTRUCCIONES:
- Se debe ejecutar en la máquina dist69 el comando: ``` X ```
- Se debe ejecutar en la máquina dist70 el comando: ``` X ``` 
- Se debe ejecutar en la máquina dist71 el comando: ``` X ``` 
- Se debe ejecutar en la máquina dist72 el comando: ``` X ``` 

Los comandos que se envían desde el informante pueden ser de 4 tipos:
- Tipo "0": Se crea un planeta, con su ciudad y el valor asociado.
- Tipo "1": SE actualiza el nombra de la ciudad.
- Tipo "2": Se actualiza el valor asociado.
- Tipo "3": Se elimina el planeta.
Esto servirá para el log en los servidores fulcrum y anotar que está ocurriendo.

## CONSIDERACIONES
- Se asume que tanto los informantes, como leia son inteligentes y seguirán las instrucciones que aparecen en pantalla.
- Se asumirá que los informantes al agregar una ciudad no agregarán ciudades respetidas.
- Al momento de actualizar un valor (ya sea ciudad o valor), se asume que se indicará un lugar existente.
- El reloj está pensado como un reloj universal, donde cada componente corresponde al Servidor Fulcrum al cual se le están realizando modificaciones.
- La consistencia read your writes se hace en base a lo que envía el informante y en base a lo que recibe de vuelta, donde deben ser los mismos datos.

RECORDAR BORRAR LO SIGUIENTE:

- Para actualizar el archivo proto se debe ejecutar el siguiente comando: protoc --go_out=plugins=grpc:chat chat.proto
- Ejecutar el broker con el comando: go run brokerMosEisley.go
- Ejecutar los informantes con el comando: go run informante.go