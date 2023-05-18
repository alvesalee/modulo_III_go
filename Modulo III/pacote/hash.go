//hash são usadas em tudo em programação, principalmente em dados e detectação
//aceita um conjunto de dados e reduz a um tamanho fisico menor
//em go são divididas as criptografadas e não criptografadas
//Lista de não crip : ADLER32 , CRC32 , CRC64
// Nesses codigos vamos criar um hash e por nossos dados nele

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	//Primeiro de tudo: Criar a hash
	h := crc32.NewIEEE()
	//escrever os dados do hash
	h.Write([]byte("Código com pacote hash"))
	//Calcular o hash
	v := h.Sum32()
	//Criar outra variavel para o valor hash
	fmt.Println(v)

}
