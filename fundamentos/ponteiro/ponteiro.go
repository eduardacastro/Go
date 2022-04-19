package main

import "fmt"

func main() {
	i := 1

	//Go não tem aritmética de ponteiros

	var p *int = nil // nil = null

	p = &i //pegando o endereço da variavél
	*p++   // desreferencindo (pegando o valor)

	i++

	fmt.Println(p, *p, i, &i)

	/* & : lugar de memoria
	 *ponteiro : o que tá dentro da variavel que o ponteiro aponta
	 */

}
