package main

import (
	"fmt"
)

func main() {
	var emptyInterfaceVar interface{}
	describe(emptyInterfaceVar)

	emptyInterfaceVar = 42
	describe(emptyInterfaceVar)

	emptyInterfaceVar = "HELLO PES"
	describe(emptyInterfaceVar)
}

func describe(emptyInterfaceVar interface{}) {
	fmt.Printf("(%v, %T)\n", emptyInterfaceVar, emptyInterfaceVar)
}
