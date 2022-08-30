package main

import (
	"fmt"
)

func main() {
	//Maps es un diccionario como podria ser en python
	//Make nos permite crear diccionarios y otras variables
	m := make(map[string]int)

	m["Jose"] = 14
	m["pepito"] = 12

	fmt.Println(m)

	//Para recorrer un diccionario
	for i, v := range m {
		fmt.Println(i, v)
	}

	//encontrar un valor
	value, ok := m["Jose"] //el ok es un booleano que nos confirma si esa llave existe o no
	fmt.Println(value, ok)

}
