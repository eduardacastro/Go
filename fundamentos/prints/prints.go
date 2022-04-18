package main

import "fmt"

func main() {

	// Print: printa na msm linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println()
	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.141516

	// Transformando o inteiro em string
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	//Outras soluções que não precisam da concatenação
	fmt.Println("O valor de x é", x, "!!!!!")

	fmt.Printf("O valor de x é %f", x)
	fmt.Println()
	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	//Usando v, impre quase todos de forma correta
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
