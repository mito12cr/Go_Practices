package main

func main() {
	var numero int
	numero = 30
	println(numero)

	nombre := "jaime"
	println(nombre)

	nombre2 := "felipe"
	println(nombre2)

	nombre2, nombre = nombre, nombre2
	println(nombre, nombre2)

	nombre3, nombre := "Maria", "Rafael"
	println(nombre3, nombre)
}
