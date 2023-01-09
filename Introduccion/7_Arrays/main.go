package main

import "fmt"

func main() {
	//arrays
	var numeros [5]int //asi se define un array de 5 posiciones
	fmt.Println(numeros)

	//Insertamos valores de acuerdo al indice correspondiente, programcion se empieza a contar desde 0
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println(numeros)
	fmt.Println(numeros[4]) //Imprimimos el valor que se encuentra en la posicion 4 del arreglo numeros

	//Arrays con valores
	nombres := [3]string{"Alex", "Roel", "Juan"}

	fmt.Println(nombres)

	colores := [...]string{
		"Rojo",
		"Verde",
		"Negro",
		"Azul",
	}

	//la funcion LEN se usa para saber la longitud de cualquier array
	fmt.Println(colores, len(colores))

	//Indices definidos o llave valor
	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euro",
	}

	fmt.Println(monedas, len(monedas))

	//sub Array
	subMoneda := monedas[3:]
	fmt.Println(subMoneda)

}
