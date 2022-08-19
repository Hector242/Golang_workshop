package main

import "fmt"

func main() {
	//Declaracion de constantes
	const Hellomessage string = "Hello"
	const Worldmessage string = "World"

	//Println
	fmt.Println(Hellomessage, Worldmessage)
	fmt.Println(Hellomessage, Worldmessage)

	//Printf
	//Sirve para introducir variables o constantes en el mensaje a imprimir en pantalla
	const name string = "Hector"
	const lastname_count int = 2

	fmt.Printf("%s have more than %d lastnames\n", name, lastname_count)
	fmt.Printf("%v have more than %v lastnames\n", name, lastname_count) // %v la usamos cuando no sepamos el tipo de dato

	// Sprintf
	//Sprintf sirve para guardar un print y que no se muestre por consola
	var message string = fmt.Sprintf("%s have more than %d lastnames", name, lastname_count)
	fmt.Println(message)

	//data type
	//El %T nos dice el tipo de dato
	fmt.Printf("hellomessage: %T\n", Hellomessage)
	fmt.Printf("lastname count: %T\n", lastname_count)

}
