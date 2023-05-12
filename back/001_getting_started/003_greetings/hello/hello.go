package main

import "fmt"
import "greetings/greetings"

func main() {
	message := greetings.Hello("PES")
	fmt.Println(message)
}
