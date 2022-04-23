package main

import "math"

// Inicializando com letra maiuscula é PUBLICO (visibilidade dentro e fora do pacote)
// Inicializando com letra minuscula é PRIVADO dentro do pacote (visivel dentro do pacote)

// Ponto - gera algo publico
// ponto ou _Ponto - geraré algo privado

// Ponto que representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

func Distacia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
