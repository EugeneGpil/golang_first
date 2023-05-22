package main

import (
	"fmt"
	"math"
)

type MyInterface interface {
	MyMethod()
}

type MyType struct {
	MyString string
}

func (t *MyType) MyMethod() {
	fmt.Println(t.MyString)
}

type MyFloat float64

func (f MyFloat) MyMethod() {
	fmt.Println(f)
}

func main() {
	var interfaceVariable MyInterface

	interfaceVariable = &MyType{"PES"}
	describe(interfaceVariable)
	interfaceVariable.MyMethod()

	interfaceVariable = MyFloat(math.Pi)
	describe(interfaceVariable)
	interfaceVariable.MyMethod()
}

func describe(interfaceVariable MyInterface) {
	fmt.Printf("(%v, %T)\n", interfaceVariable, interfaceVariable)
}
