package main

import (
	"fmt"
)

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (pro produto) toString() string {
	spreco := fmt.Sprintf("%.2f", pro.preco)
	return pro.nome + " " + spreco
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Robertp", "Viana"}
	//fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Lápiz", 6.982}
	//fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça", 89.99}
	imprimir(coisa)
}
