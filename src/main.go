package main

import (
	"fmt"
)

func main() {
	// Array
	//son Inmutables
	var array [4]int //se le indica la cantidad de valores a guardar
	array[0] = 1
	array[1] = 2
	//len te da la longitud del array
	//cap te da su capacidad maxima
	fmt.Println(array, len(array), cap(array))

	//slices
	//son Mutables
	slice := []int{0, 1, 2, 4, 5} //no se le agrega la cantidad de valores a guardar
	fmt.Println(slice, len(slice), cap(slice))

	//metodos en el slice
	fmt.Println(slice[0])   //imprime el valor del indice 0
	fmt.Println(slice[:3])  //imprime los valores de los indices del 0 al 3
	fmt.Println(slice[2:4]) //imprime los valores de los indices del 2 al 4
	fmt.Println(slice[4:])  //imprime los valores de los indices del 4 al ultimo

	//append
	slice = append(slice, 6) //agrega el numero 6 al slice
	fmt.Println(slice)

	//append nueva lista
	newSlice := []int{7, 8, 9}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
