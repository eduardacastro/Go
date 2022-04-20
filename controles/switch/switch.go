package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6:
		return "C"
	case 5, 4, 3, 2, 1:
		return "D"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(9.7))
	fmt.Println(notaParaConceito(6.2))
	fmt.Println(notaParaConceito(3.0))
	fmt.Println(notaParaConceito(0))
}
