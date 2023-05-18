// O pacote lista, implementa uma lista duplamente encadeada
// Ex: Uma lista de numeros de lugares diferentes, ele "liga" um a outro, conexões

package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

//criar uma lista, puxar os numeros de outras variaveis
