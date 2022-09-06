package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println("la marca es: ", myPc.brand)
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func main() {

	a := 50
	b := &a

	fmt.Println("La direccion de memoria es: ", b)
	fmt.Println("El valor en esa direccion es: ", *b)

	*b = 100
	fmt.Println("Nuevo valor de la variable a es: ", a)

	// ejemplo de structs con funciones
	var myPc pc
	myPc.brand = "asus"
	myPc.disk = 100
	myPc.ram = 32

	fmt.Println("My pc tiene: ", myPc)

	myPc.ping()

	// usando punteros con los struct en funciones
	fmt.Println("Estado actual de mi pc: ", myPc)

	myPc.duplicateRam()
	fmt.Println("Ram duplicada: ", myPc)
}
