package main

import "fmt"

//función que calcula suma
func suma(a, b int) int {
	return a + b
}

//función que calcula resta
func resta(a, b int) int {
	return a - b
}

//función que calcula multiplicacion
func Multiplicacion(a, b int) int {
	return a * b
}

//función que calcula cociente
func cociente(a, b int) int {
	return a / b
}

//Función que calcula residuo
func residuo(a, b int) int {
	return a % b
}

//Función principal
func main() {

	//Variables
	var a, b, suma1, resta1, multiplicacion, cociente1, residuo1 int

	//Entrada  de datos
	fmt.Print("Ingrese N01: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese N02: ")
	fmt.Scanln(&b)

	//Llamar la función suma
	suma1 = suma(a, b)

	//Llamar la función resta
	resta1 = resta(a, b)

	//Llamar la función multiplicacion
	multiplicacion = Multiplicacion(a, b)

	//Llamar la función cociente
	cociente1 = cociente(a, b)

	//LLamar la función residuo
	residuo1 = residuo(a, b)

	//Salida de datos
	fmt.Println("suma:", suma1)
	fmt.Println("resta:", resta1)
	fmt.Println("multiplicacion:", multiplicacion)
	fmt.Println("Cociente:", cociente1)
	fmt.Println("Residuo:", residuo1)
}
