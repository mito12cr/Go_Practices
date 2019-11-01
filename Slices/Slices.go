package main

import "fmt"

func main() {

	//Slices

	//Declaraciones de Slices
	var j []int
	fmt.Println("slice j", j)

	//Ejemplo 2
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)

	//declaring slices using "Make" indicating the length
	y := make([]int, 10, 23)
	fmt.Println("slice y:", y)
	fmt.Println("Longitud:", len(y))
	fmt.Println("capacidad de y: ", cap(y))

	fmt.Println("********************************************************************************")

	//Un solo array con 10 elementos
	x := [15]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)

	fmt.Println("********************************************************************************")

	//Doble
	var a, b []int
	fmt.Println("Slice o: ", a)
	fmt.Println("Slice p: ", b)

	a = x[2:8]
	fmt.Println("Slice a: ", a)

	b = x[8:12]
	fmt.Println("Slice b: ", b)
}
