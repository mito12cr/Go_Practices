package main

import (
	"fmt"
)

func main() {

	//Imprimir el numero del dia de la semana que se digite

	// var dia int

	// fmt.Println("Digite el número del dia de la semana")

	// fmt.Scanln(&dia)

	// switch dia {
	// case 1:
	// 	fmt.Println("usted digitó Lunes.")
	// case 2:
	// 	fmt.Println("usted digitó Martes.")
	// case 3:
	// 	fmt.Println("usted digitó Miercoles.")
	// case 4:
	// 	fmt.Println("usted digitó Jueves.")
	// case 5:
	// 	fmt.Println("usted digitó Viernes.")
	// case 6:
	// 	fmt.Println("usted digitó Sabado.")
	// case 7:
	// 	fmt.Println("usted digitó Domingo.")
	// default:
	// 	fmt.Println("Número Inválido.")
	// }

	numero := 15
	switch {
	case numero > 40:
		fmt.Println("Es mayor que 40")
		fallthrough
	case numero > 20:
		fmt.Println("Es mayor que 20")
		fallthrough
	case numero > 10:
		fmt.Println("Es mayor que 10")
	default:
		fmt.Println("No es mayor que ninguno")

	}
}
