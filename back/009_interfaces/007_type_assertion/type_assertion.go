package main

import (
	"fmt"
)

func main() {
	var emptyInterfaceVar interface{} = "HELLO PES"

	stringVar := emptyInterfaceVar.(string)
	fmt.Println(stringVar)

	stringVar, isConversionSuccessful := emptyInterfaceVar.(string)
	fmt.Println(stringVar, isConversionSuccessful)
	
	floatVar, isConversionSuccessful := emptyInterfaceVar.(float64)
	fmt.Println(floatVar, isConversionSuccessful)

	// floatVar = emptyInterfaceVar.(float64) // panic
	// fmt.Println(floatVar)
}
