package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bites long\n", v, len(v))
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}

func main() {
	do(42)
	do("HELLO PES")
	do(true)
}
