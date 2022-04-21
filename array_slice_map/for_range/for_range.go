package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //copilador conta o numeros de elementos

	for i, numero := range numeros { // i = indice/ numero = numero dentro daquele indice
		fmt.Printf("%d) %d \n", i+1, numero)

	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
