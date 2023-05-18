package main

import (
	"io/ioutil"
	"log"
)

func main() {

	message := []byte("Teste Go!")
	err := ioutil.WriteFile("Teste", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// Esse código cria um novo arquivo, com as instruções que estão na fatia, nome e o texto logo após.
