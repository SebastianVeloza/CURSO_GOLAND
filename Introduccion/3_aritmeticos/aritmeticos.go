package main

import "fmt"

func main() {
	//Definimos dos variables
	numero1 := 10
	numero2 := 50

	//Suma
	resultado := numero1 + numero2
	fmt.Println("Suma: ", resultado)

	//resta

	resultado = numero1 - numero2 //los dos puntos solo se usan para definir la variable, si la queremos modificar con el "= "
	fmt.Println("Resta: ", resultado)

	//Multiplicacion

	resultado = numero1 * numero2
	fmt.Println("Multiplicacion: ", resultado)

	//Division

	resultado = numero1 / numero2

	fmt.Println("Division: ", resultado)

	//modulo

	resultado = numero1 % numero2
	fmt.Println("Modulo: ", resultado)
}
