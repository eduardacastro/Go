package main

import (
	"fmt"
	"time"
)

func fale(pessoa, text string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, text, i+1)
	}
}

func main() {
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("Julia", "Só posso falar depois de voce.", 1)

	go fale("Maria", "Ei...", 500)
	go fale("Julia", "Opa...", 500)

	time.Sleep(time.Second * 5)
	fmt.Println("Fim")
}
