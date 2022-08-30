package main

import (
	"fmt"
)

type car struct {
	brand string
	year  int
}

func main() {
	// first way to instance a struct
	myCar := car{brand: "Toyota", year: 2022}
	fmt.Println(myCar)

	// other way to instance a struct
	var secondCar car
	secondCar.brand = "Ford"
	secondCar.year = 2001
	fmt.Println(secondCar)
}
