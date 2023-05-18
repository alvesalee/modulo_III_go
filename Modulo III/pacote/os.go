package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile("AulaGo.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

//Código abre outro arquivo para armazenar algo
//OpenFile é a chamada aberta generalizada; a maioria dos usuários usará Abrir ou Criar. Ele abre o arquivo nomeado com o sinalizador especificado (O_RDONLY etc.). Se o arquivo não existir e o sinalizador O_CREATE for passado, ele será criado com o modo perm (antes de umask). Se for bem-sucedido, os métodos no arquivo retornado podem ser usados ​​para E/S. Se houver algum erro, ele será do tipo *PathError.
