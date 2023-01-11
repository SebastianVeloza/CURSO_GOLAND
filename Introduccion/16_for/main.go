package main

func main() {
	//Nota: En GOLAND NO existe en while pero si se puede usar el for como un while

	//Bucle infinito
	/*for {
		fmt.Println("Hola Mundo")
	}*/

	//Bucle tipo while

	numeros := 12455458
	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}
}
