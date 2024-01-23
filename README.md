# Saludos de go

Hola este paquete te proporciona una manera sencilla de saludos personalizados en lenguaje de GO.

# Instalacion

Ejecuta el siguiente comando para instalar el paquete:

```bash

go get -u github.com/HeavysetSine/greetings

```

# Ejemplo de Uso

Aqui te proporciono un ejemplo de como imprementar y usar el paquete en tu codigo:

```go
package main
import (
	"fmt"
	"log"
	"github.com/HeavysetSine/greetings"
)

func main() {

	mensaje, error := greetings.Hello("Jaime")

	if error != nil {
		log.Fatal("Ocurrio un error",error) ///Imprimiremos el error
        return
	}
	fmt.Println(mensaje) //evidenciamos el mensaje

}
```
