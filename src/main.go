package main

import (
	"curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	var mycar mypackage.CarPublic
	mycar.Brand = "ferrary"
	mycar.Year = 2022
	fmt.Println(mycar)

	mypackage.PrintMessagePublic()
}
