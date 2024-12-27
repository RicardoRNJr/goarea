package goarea

import "math"

//PI é a relação entre o perimetro e o diametro de uma circunferencia
const Pi = 3.1416

//Circ calcula a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Nao é vissivel, teste
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
