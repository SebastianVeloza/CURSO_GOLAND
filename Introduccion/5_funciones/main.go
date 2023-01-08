package main

import "fmt"

//Funcion sin parametros
func saludar() {
	fmt.Println("Hola Mundo")
}

//funcion con parametros
func saludar2(nombre string) {
	fmt.Println("Hola,", nombre)
}

//funcion con retorno
func sumar(a, b int) int {
	return a + b
}

func main() {
	//Llamar las funciones
	saludar()
	saludar2("Alex")
	resultado := sumar(10, 20)
	fmt.Println("La suma es: ", resultado)

	//******************************************************************************
	//Ahora recibiendo los parametros de la suma con el scan
	//Definimos las variabels
	var a, b, suma int

	//Entrada  de datos
	fmt.Print("Ingrese N01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese N02: ")
	fmt.Scanln(&b)

	//Llamar la función sumar nuevamente como es una funcion la podemos usar ilimitadamente
	suma = sumar(a, b)

	//Salida de datos
	fmt.Println("La suma: ", a, " + ", b, " = ", suma)
}
