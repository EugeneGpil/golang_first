package main

import (
	"fmt"
)

type MyInterface interface {
	MyMethod()
}

func main() {
	var myInterfaceVar MyInterface

	describe(myInterfaceVar)
	myInterfaceVar.MyMethod()
}

func describe(myInterfaceVar MyInterface) {
	fmt.Printf("(%v, %T)\n", myInterfaceVar, myInterfaceVar)
}
