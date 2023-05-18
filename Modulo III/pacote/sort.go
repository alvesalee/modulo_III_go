// Ordenar dados arbitrarios
// pacote sort.Interface
//Sort.interface ira ordenar 3 metodos (Len, less, swap)

// Package sort fornece primitivas para classificar fatias e coleções definidas pelo usuário.

package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome  string
	Idade int
}

type ParaNome []Dados

func (ps ParaNome) Len() int {
	return len(ps)
}
func (ps ParaNome) Less(i, j int) bool {
	return ps[i].Nome < ps[j].Nome
} //Less = Vai ver se o item da posição i é menor que o item na posição j

func (ps ParaNome) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

//swap = Ele vai trocar os itens

func main() {

	criancas := []Dados{
		{"Alexandre", 15},
		{"Isabela", 12},
	}
	sort.Sort(ParaNome(criancas))
	fmt.Println(criancas)
}
