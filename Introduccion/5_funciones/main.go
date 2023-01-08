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
	saludar()
	saludar2("Alex")
	resultado := sumar(10, 20)
	fmt.Println("La suma es: ", resultado)

	//Ahora recibiendo los parametros de la suma con el scan
}
