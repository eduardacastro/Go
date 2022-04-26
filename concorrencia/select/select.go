package main

import (
	"fmt"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := htmltitulo.Titulo(url1)
	c2 := htmltitulo.Titulo(url2)
	c3 := htmltitulo.Titulo(url3)

	//estrutura de controle especifica para concorrencia
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(time.Second * 100):
		return "Todos perderam!"
	}
}
func main() {
	campeao := oMaisRapido("https://youtube.com", "http://www.google.com", "http://www.amazon.com")
	fmt.Println(campeao)
}
