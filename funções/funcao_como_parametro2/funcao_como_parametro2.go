package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados:")

	for _, aprovado := range aprovados {
		fmt.Println(aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Ana", "João", "Luis"}
	imprimirAprovados(aprovados...)
}
