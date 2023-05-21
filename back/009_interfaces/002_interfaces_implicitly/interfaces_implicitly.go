package main

import (
	"fmt"
)

type Interface interface {
	Method()
}

type Type struct {
	String string
}

func (typeVar Type) Method() {
	fmt.Println(typeVar.String)
}

func main() {
	var interfaceVar Interface = Type{"PES"}
	interfaceVar.Method()
}
