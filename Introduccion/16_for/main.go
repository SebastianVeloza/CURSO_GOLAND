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

}
