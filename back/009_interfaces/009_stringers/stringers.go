package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"PES PES", 14}
	z := Person{"PES PES PES", 18}

	fmt.Println(a)
	fmt.Println(z)
}
