package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string

	// Volteamos el texto
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i]) //pasamos de numero ASCII a texto
	}

	//imprimimos palabra inversa
	fmt.Println("Palabra inversa: ", textReverse)

	//Convertimos todas las palabras en minuscula para la comparacion
	var text_lowercase string = strings.ToLower(text)
	var textReverse_lowercase string = strings.ToLower(textReverse)

	//Revisamos palabra inversa
	if text_lowercase == textReverse_lowercase {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	// Imprimir  slice con indice
	for i, value := range slice {
		fmt.Println(i, value)
	}
	// Imprimir  slice sin indice
	for _, value := range slice {
		fmt.Println(value)
	}
	// Imprimir  indice
	for value := range slice {
		fmt.Println(value)
	}

	//Palindromo
	var word string

	fmt.Printf("\n")
	fmt.Printf("Ingrese una palabra: ")
	fmt.Scanf("%s", &word)

	fmt.Printf("\n")
	isPalindromo(word)
}
