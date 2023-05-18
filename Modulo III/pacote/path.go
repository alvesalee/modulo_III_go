package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

// Implementa rotinas de utilitario para manipular caminhos de nome de arquivo
// Percorre toda a aula pacote e devolvendo os arquivos da pasta
