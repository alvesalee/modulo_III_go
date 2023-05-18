// Exemplo de codigo pegado na documentação da linguagem
// C´´odigo para não usar o pacote fmt para colocar alguma mensagem para aparecer
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if _, err := io.WriteString(os.Stdout, "Olá pexualll"); err != nil {
		log.Fatal(err)
	}
}
