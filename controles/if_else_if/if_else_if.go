package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 18 {
		return "A"
	} else if nota >= 7 && nota < 9 {
		return "B"
	} else if nota >= 6 && nota < 7 {
		return "C"
	} else {
		return "D"
	}
}
func main() {
	nota := notaParaConceito(8)
	fmt.Println(nota)
}
