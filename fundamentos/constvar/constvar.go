package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.141592653589
	var raio = 3.2 //tipo (float64)  inferido pelo copilador

	//forma reduzida de criar um var

	area := PI * math.Pow(raio, 2)
	fmt.Println("Area da circunferencia: ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := false, 6, "ola"

	fmt.Println(g, h, i)
}
