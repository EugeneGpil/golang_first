package main

import "fmt"

func fibonacci() func() int {
	iteration := 0

	prev_fibonacci := 0
	fibonacci := 1
	return func() int {
		if iteration == 0 {
			iteration = 1

			return 0
		}

		if iteration == 1 {
			iteration = 2

			return 1
		}
		
		future_prev_fibonacci := fibonacci
		fibonacci = fibonacci + prev_fibonacci 
		prev_fibonacci = future_prev_fibonacci

		return fibonacci
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
