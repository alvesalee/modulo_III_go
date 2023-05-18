// Goroutines: é uma função capaz de executar de modo concorrente com outras funções
// Palavra reservada para chamar o goroutines é "GO"
package main

import "fmt"

func funcao(numero int) {
	for i := 0; i < 10; i++ {
		fmt.Println(numero, ":", i)
	}
}

func main() {
	go funcao(0)
	var escreva string
	fmt.Scanln(&escreva) // ScanLn: Scanln é semelhante a Scan, mas para de escanear em uma nova linha e após o item final deve haver uma nova linha ou EOF.

}
