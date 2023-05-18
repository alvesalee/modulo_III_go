// Select: funciona como um switch para canais
// Select Permite que voce aguarde operação de varios canais
// Combinar goroutines e canais com select é um recurso poderoso do go

package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() { // Recebemos os valores "UM" e depois o valor "DOIS" como o esperado
		time.Sleep(1 * time.Second) // Observe que o tempo total é de apenas
		// ~ 2 segundos
		// Pois o 1 e 2 são sleep executados simultaneamente
		canal1 <- "um"

	}()

	go func() {
		time.Sleep(2 * time.Second)
		canal2 <- "Dois"
	}()

	for i := 0; i < 2; i++ {
		select { // Usaremos o select para aguardar esses dois valores simultaneamente, imprimindo cada um a medida que chegam
		case msg1 := <-canal1:
			fmt.Println("Receba", msg1)
		case msg2 := <-canal2:
			fmt.Println("Receba", msg2)
		}

	}
}
