package main

import (
	"fmt"
)

func main() {

	//Uso de funcion recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ups!, al parecer el programa no finalizo de forma correcta")
		}
	}()

}
