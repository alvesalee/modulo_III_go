// Pacote erro = implementa funções para anipular erros
// cria funcoes cujo o conteudo é criar uma mensagem de texto

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"O arquivo do sistema desapareceu",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}

//Esse código vai mostrar o horario, data e aparecer a mensagem que lhe foi escrita para falar
//o Arquivo do sistema desapareceu
