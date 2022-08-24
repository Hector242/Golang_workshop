package main

import (
	"fmt"
)

func main() {

	switch module := 5 % 2; module {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//without condition
	value := 200
	switch {
	case value > 100:
		fmt.Println("es mayor a 100")
	case value == 100:
		fmt.Println("es igual a 100")
	default:
		fmt.Println("Es menor a 100")
	}
}
