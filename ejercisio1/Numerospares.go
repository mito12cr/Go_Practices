package main

import (
	"fmt"
)

func main() {

	//contador de numeros pares

	encabezado := `
	*****************************
	Contenedor de Numeros Impares
	*****************************
	`

	//Impresion del emcabezado
	fmt.Println(encabezado)
	//Se pide el primer numero
	fmt.Println("Digite el Primer Numero")
	//El numero se almacena en la variable numero1
	var numero1 int
	//Leemos el numero digitado y lo guardamos en la variable numero1
	fmt.Scanln(&numero1)
	//Mismo procedimiento para el numero 2
	fmt.Println("Digite el segundo numero")
	var numero2 int
	fmt.Scanln(&numero2)
	//Declaramos una variable contador para guardar la cantidad de  numeros
	var contador int
	//realizamos un bucle tomando como inicio y fin de los numeros digitados
	for i := numero1; i <= numero2; i++ {
		//evaluamos si el numero es impar
		if i%2 != 0 {
			//incrementamos el valor de la variable contador 1
			contador++
			//imprimimos el numero impar
			fmt.Println("%d es un numero impar\n", i)
		}
	}

	encabezado = `
	*****************************
	Contenedor de Numeros Impares
	*****************************
	`
	//imprimimos el emcabezado
	fmt.Println(encabezado)
	//imprimimos los resultados
	fmt.Println("Entre el %d y el %d hay %d numeros impares.", numero1, numero2, contador)
}
