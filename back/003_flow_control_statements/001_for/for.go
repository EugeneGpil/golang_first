package main

import "fmt"

func func_001() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func func_002() {
	sum := 1
	// init and post statements are optional
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func func_003() {
	sum := 1
	// and at this point can drop semicolons
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func func_004() {
	// infinit loop
	for {
		
	}
}

func main() {
	func_001()
	func_002()
	func_003()
	// func_004()
}
