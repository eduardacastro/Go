//Map Aninhado

package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15546.78,
			"Guga Pereira":   7689.9,
		},
		"J": {
			"José Junior": 7839.8,
			"João Luis":   6875.9,
		},
		"P": {
			"Patricia Luisa":  70876.9,
			"Petunia Phelipa": 5769.9,
		},
	}

	delete(funcPorLetra, "J")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}

	for letra, funcs := range funcPorLetra {
		fmt.Println("\n", letra, ":")
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
