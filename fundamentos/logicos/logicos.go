package main

import "fmt"

func compras(trab_terca, trab_quinta bool) (bool, bool, bool) {

	comprarTV50 := trab_terca && trab_quinta
	comprarTV32 := trab_terca != trab_quinta //ou exclusivo
	comprarSorvete := trab_terca || trab_quinta

	return comprarTV50, comprarTV32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
