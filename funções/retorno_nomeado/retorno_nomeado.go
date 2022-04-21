package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {

	segundo, primeiro = p2, p1

	return //retorno limpo
}

func main() {

	segundo, primeiro := troca(3, 7)
	fmt.Println(segundo, primeiro)

}
