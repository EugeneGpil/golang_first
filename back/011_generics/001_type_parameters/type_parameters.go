package main

import (
	"fmt"
)

// Index retruns position of element, or -1 if not found
func Index[T comparable](whereToSearch []T, elementValue T) int {
	for index, value := range whereToSearch {
		if value == elementValue {
			return index
		}
	}

	return -1
}

func main() {
	sliceOfInts := []int{10, 20, 15, -10}
	fmt.Println(Index(sliceOfInts, 20))

	sliceOfStrings := []string{"HELLO", "PES"}
	fmt.Println(Index(sliceOfStrings, "PES"))
}
