package main

import "fmt"

func somar(a int, b int) int {
	return (a + b)
}

func imprimir(a int, b int, valor int) {
	fmt.Printf("%v + %v = %v", a, b, valor)
}

func main() {
	a := 3
	b := 6

	resultado := somar(a, b)
	imprimir(a, b, resultado)
}
