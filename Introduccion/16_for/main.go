package main

import "fmt"

func main() {
	//Nota: En GOLAND NO existe en while pero si se puede usar el for como un while
	/*for {
		fmt.Println("Hola Mundo")
	}*/

	//Bucle tipo while
	/**
	numeros := 12455458
	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}
	fmt.Println("Cantidad de digitos es: ", c)
	*/

	//Bucle for

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//Continue y break

	for i := 0; i <= 10; i++ {

		if i == 5 {
			fmt.Println("Salta la iteración")
			continue
		}

		if i == 8 {
			fmt.Println("Romper con bucle ")
			break
		}
		fmt.Println(i)
	}

	//Bucle For each

	nombres := [...]string{"Alex", "Roel", "Juan"}

	/**
	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}*/

	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}

}
