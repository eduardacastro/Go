package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado
	fmt.Println("Teste " + string(97))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para inteiro
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// apenas o valor 1 e o "true" vão fazer com que imprimo o "Verdadeiro"
	b, _ := strconv.ParseBool("1")
	if b {
		fmt.Println("Verdadeiro")
	}
}
