package main

import "fmt"

func main() {

	funcESalarios := map[string]float64{
		"José Silva":     11456.89,
		"Gabriela Thais": 4598.98,
		"Ana Joaquina":   40188.9,
		"Maria Marta":    10987.9,
		"Ana Gabriela":   3879.7,
		"Mario João":     7650.0,
	}

	funcESalarios["Naiomi Nami"] = 6789.0
	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}

	for _, salario := range funcESalarios {
		fmt.Println(salario)
	}

	for salario := range funcESalarios {
		fmt.Println(salario)
	}
}
