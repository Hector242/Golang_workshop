package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi", pi2)

	//Declaracion de variables enteras. 3 formas:

	base := 10 //los dos puntos indican que se esta declarando la variable por primera vez
	var altura int = 5
	var area int

	area = (base * altura) / 2

	fmt.Println("El area es: ", area)
}
