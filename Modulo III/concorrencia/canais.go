// Canais: forma de duas goroutines se comunicarem e sincronizarem a sua execução

package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" // ( <- ) Sinal usado para enviar e receber mensagem pelo canal

	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)

	}
}

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)

}

//Código vai imprimir a palavra "ping", num loop infinito, pois não falou quando deve parar nas funções
//para parar é so apertar o enter
//time: vai esperar o tempo que foi pedido no codigo para colocar o valor na proxima linha
