// Cripto são exemplos criptografados
//vamos usar uma muito comum o sha-1

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()
	h.Write([]byte("Código com pacote hash"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}

//Mesma frase e com resultado bem diferente
