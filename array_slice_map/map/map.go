package main

import "fmt"

func main() {

	// var aprovados map[int]string
	// mapas devem ser inicializados

	// estrutura chave/valor
	/*

		NomeDoMap := map[tipo da chave do map]tipo do valor do map
	*/

	aprovados := make(map[int]string)

	aprovados[123456789] = "Ana"
	aprovados[123456799] = "Joao"
	aprovados[123456788] = "Maria"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123456789])

	delete(aprovados, 123456789)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
}
