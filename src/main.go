package main

import (
	"fmt"
)

func main() {

	//defer.
	//Defer es el keyword que ejecuta  la funcion antes que el programa muera
	defer fmt.Println("Hello")
	fmt.Println("world")

	//continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		//permite continuar un ciclo a pesar que se de una condicion
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		//permite parar un ciclo si se da una condicion
		if i == 8 {
			fmt.Println("Es break")
			break
		}
	}
}
