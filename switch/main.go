package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "Segunda"
	case 2:
		return "Terca"
	case 3:
		return "Quarta"
	case 4:
		return "Quinta"
	case 5:
		return "Sexta"
	case 6:
		return "Sabado"
	default:
		return "Domingo"
	}
}

func main() {
	fmt.Println("Anotações Switch")

	calendarioToday := diaSemana(3)
	fmt.Println(calendarioToday)

}
