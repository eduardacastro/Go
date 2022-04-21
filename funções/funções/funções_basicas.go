package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %v %v\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("\nF4: %v %v", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Parametro 1", "Parametro 2")

	r3, r4 := f3(), f4("Parametro 1", "Parametro 2")
	fmt.Println(r3, r4)

	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)

}
