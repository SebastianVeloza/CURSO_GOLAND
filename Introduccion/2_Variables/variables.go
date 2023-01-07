package main

import "fmt"

func main() {

	/*las variables se pueden definir de dos maneras una tipo JavaScript
	var "nombre de la variable" "tipo de dato"

	Y las tipo Pyhton
	"nombre de la variable := "lo que queramos almacenar cualquier tipo de dato"*/

	var cadena string
	var entero int
	var boleano bool
	var flotante float64
	fmt.Println("Variables definidas \n", cadena, entero, boleano, flotante)

	//Definiendo varible

	cadena = "Canal de twich Matockf"
	entero = 1
	boleano = true
	flotante = 1.9
	fmt.Println("Variables almacenadas \n", cadena, entero, boleano, flotante)

	//Variables tipo python
	cadena2 := "sigueme"
	entero2 := 2
	boleano2 := false
	flotante2 := 1.2

	fmt.Println("Variables almacenadas tipo python \n", cadena2, entero2, boleano2, flotante2)

	//Variable constante es inmutable
	const version = "NOS VEMOS EN LA PROXIMA CLASE"
	fmt.Println(version)

}
