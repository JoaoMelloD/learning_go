package main

import (
	"fmt"
	"math"
)

// So tem assinaturas de métodos que falam como esses métodos devem ser
type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f", f.area())
}

func main() {
	r := retangulo{10, 20}
	escreverArea(r)
}
