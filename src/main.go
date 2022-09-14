package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) String() string {
	return fmt.Sprintf("my pc has %d Gb of ram, %d Tb of disk and is a %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {

	var myPc pc
	myPc.brand = "asus"
	myPc.disk = 100
	myPc.ram = 32

	fmt.Println(myPc)
}
