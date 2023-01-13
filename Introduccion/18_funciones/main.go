package main

import "fmt"

//Función variádica -
func sumar(nombre string, numeros ...int) (string, int) {

	mensaje := fmt.Sprintf("La suma de %s es: ", nombre)
	var total int
	for _, num := range numeros {
		total += num
	}

	//Retornar multiples valores
	return mensaje, total
}

//Función recursica
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {

	//Variadica
	mensaje, result := sumar("Alex", 10, 20, 40, 70, 60)
	fmt.Println(mensaje, result)

	//Recursiva
	fmt.Println(factorial(5))
}
