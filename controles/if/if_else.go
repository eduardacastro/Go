package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Passou!", nota)
	} else {
		fmt.Println("Não passou", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.9)
	imprimirResultado(9.8)
}
