package main

import "fmt"

func main() {
	p1 := Ponto{2, 4}
	p2 := Ponto{5, 8}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distacia(p1, p2))
}
