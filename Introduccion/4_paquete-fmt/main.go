package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	//Para mostrar en pantalla
	fmt.Print(hola)    //Imprime en pantalla sin salto de linea
	fmt.Println(mundo) //Imprime en pantalla con salto de linea

	nombre := "sebas"
	edad := 15

	//Para darle formato %S significa que es String y %d que es un entero
	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)

	//Si no sabemos que tipo de dato usamos %v que significa variable
	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	//Tambien lo podemos almacenar en una variable
	mensaje := fmt.Sprintf("Hola, %s y su edad es %d", nombre, edad)

	fmt.Println(mensaje)

	// el %T nos muestra el tipo de dato en este caso edad en un INT
	fmt.Printf("nombre: %T \n", edad)

	fmt.Print("Ingrese otro nombre: ")

	//El Scan se usa para recibir datos del usuario
	fmt.Scanln(&nombre)

	fmt.Println("Otro Nombre: ", nombre)
}
