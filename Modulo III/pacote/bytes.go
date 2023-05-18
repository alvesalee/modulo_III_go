// pacote bytes =  titulo
// Coloca letras em maiusculas
// Title trata s como bytes codificados em UTF-8 e retorna uma cópia com todas as letras Unicode que iniciam as palavras mapeadas para a caixa do título.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.Title([]byte("pacote bytes")))
}
