package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println(m["answer"])

	m["answer"] = 48
	fmt.Println(m["answer"])

	delete(m, "answer")
	fmt.Println(m["answer"])

	value, isset := m["anser"]
	fmt.Println(value, isset)

	map2 := make(map[string]string)

	value2, isset := map2["test"]
	fmt.Println(value2, isset)
}
