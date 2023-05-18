// Pacote = caixinha de funções
// função contains = procura dentro de uma string, uma string menor;
// exemplo: palavra Terra > voce pode pedir para ver se tem duas letras rr - se tiver vai dar condição (true)

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Terra", "rr")) // pedindo para ver se tem "rr" na palavra terra

	fmt.Println(strings.Contains("Alexandre", "feio"))
}
