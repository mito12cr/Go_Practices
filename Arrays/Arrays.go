package main

import (
	"fmt"
)

func main() {

	//declaracion de un array
	var x [4]int
	//asignaccion de un valor al array
	x[0] = 10
	x[1] = 100
	x[2] = 1000
	x[3] = 10000
	fmt.Println(x)

	//acceder a un indice en concreto
	fmt.Println(x[0])

	//Declarar e imicializar un array
	y := [10]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println(y)

	//Otra forma de declarar e inicializar un array
	j := [...]string{
		"jaime",
		"jose",
		"hedryan",
		"rafael",
		"jefrey",
	}
	fmt.Println(j)

	//funcion len() indica el tamaño  de un array
	f := [...]int{45, 87, 33, 15, 12}
	fmt.Println(f)

	//Funcion len()
	fmt.Println(len(f))
	fmt.Println(f[len(f)-1])
}
