package main

import "fmt"

//Hola Mundo
/*
Hola mundo
*/

func main() {
	var numero int
	numero = 400
	multiplicando := 25
	fmt.Println(numero * multiplicando)
	numero = 40
	fmt.Println(numero)
	nombre := "Alejandro"
	nombre = "Pedro"
	//nombre := "Carlos" NO SE PUEDE REASIGNAR EL TIPO DE DATO DE LA VARIABLES NI AL MISMO USAR EN ESTE CASO NOMBRE = "CARLOS"
	nombre, numero = "Rafael", 25
	nombre2 := "Ramon"
	nombre, nombre2 = nombre2, nombre
	fmt.Println(nombre, nombre2)
	nombre3, nombre := "maria", "Alejandro"
	fmt.Println(nombre, nombre3)
	var (
		dios    = "Goku"
		enemigo = "Cell"
	)
	fmt.Println("El dios es: " + dios + "\nsu enemigo es: " + enemigo)
}
