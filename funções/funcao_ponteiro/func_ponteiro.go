package main

import "fmt"

func inc1(n int) {
	n++
	fmt.Println(n)
}

//um ponteuro é representado por um *
func inc2(n *int) {
	// * é usado para desreferenciar, ou seja,
	// ter acesso ao valorno qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) //por valor
	fmt.Println(n)

	//revisaão: & usado para obter o endereço da variavél
	inc2(&n) // por referencia
	fmt.Println(n)
}
