package main

import "fmt"

func holamundoFuntion(message string) {
	fmt.Println(message)
}

func argument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnOneValue(a int) int {
	return a * 2
}

func returnTwoValues(a int) (c, d int) {
	return a, a * 2
}

func main() {
	holamundoFuntion("Hola mundo")
	argument(1, 2, "Hello")

	var value int = returnOneValue(2)
	fmt.Println("Value: ", value)

	var value1, value2 int = returnTwoValues(3)
	fmt.Printf("Value 1: %d", value1)
	fmt.Printf(" Value 2: %d", value2)
}
