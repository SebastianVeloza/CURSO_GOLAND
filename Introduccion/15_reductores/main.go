package main

import "fmt"

func main() {

	a := 2

	//a = a + 2

	//Operadores de asiganación

	//Suma en asignación
	a += 2
	fmt.Println(a)

	//Resta en asiganción
	a -= 2
	fmt.Println(a)

	//Multiplicación en asignación
	a *= 4
	fmt.Println(a)

	//División  en asignación
	a /= 2
	fmt.Println(a)

	//Modulo en asignación
	a %= 2
	fmt.Println(a)

	//	INCREMENTO Y DECREMENTO

	a = 0

	//Operador de incremento
	a++
	a++
	a++
	a++
	a++
	fmt.Println(a)

	//Operador de decremento

	a--
	a--
	a--
	a--
	fmt.Println(a)

}
