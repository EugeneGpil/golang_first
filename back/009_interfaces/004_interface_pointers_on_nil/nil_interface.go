package main

import (
	"fmt"
)

type MyInterface interface {
	MyMethod()
}

type MyStruct struct {
	MyString string
}

func (myStructVar *MyStruct) MyMethod() {
	if myStructVar == nil {
		fmt.Println("<nil>")

		return
	}

	fmt.Println(myStructVar.MyString)
}

func main() {
	var myInterfaceVariable MyInterface

	var myStructPointer *MyStruct
	myInterfaceVariable = myStructPointer
	describe(myInterfaceVariable)
	myInterfaceVariable.MyMethod()

	myInterfaceVariable = &MyStruct{"HELLO PES"}
	describe(myInterfaceVariable)
	myInterfaceVariable.MyMethod()
}

func describe(myInterfaceVariable MyInterface) {
	fmt.Printf("(%v, %T)\n", myInterfaceVariable, myInterfaceVariable)
}
