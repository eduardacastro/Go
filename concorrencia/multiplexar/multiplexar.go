package main

import (
	"fmt"

	"github.com/eduardacastro/httptitulo"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

//multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		httptitulo.Titulo("https://github.com/", "https://google.com"),
		httptitulo.Titulo("https://github.com/youtube.com", "https://youtubr.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
