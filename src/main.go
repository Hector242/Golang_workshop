package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var value1 int = 1
	var value2 int = 2

	if value1 == 1 {
		fmt.Println("El valor es ", value1)
	} else {
		fmt.Println("El valor no es 1")
	}

	//With AND
	if value1 == 1 && value2 == 2 {
		fmt.Println("El valor es 1 y 2")
	} else {
		fmt.Println("El valor no es el correcto")
	}

	//With OR
	if value1 == 0 || value2 == 2 {
		fmt.Println("El valor es 1 o 2")
	} else {
		fmt.Println("El valor no es el correcto")
	}

	//convert text to number
	//we this method we will get 2 values: value, err.
	value, err := strconv.Atoi("53") //On value we save the value and in err we will save the error
	if err != nil {                  //if error does not existed, it will be nil
		log.Fatal(err) //will print the error
	}
	fmt.Println("Value: ", value)
}
