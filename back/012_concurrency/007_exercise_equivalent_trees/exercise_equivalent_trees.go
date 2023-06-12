package main

import (
	"sort"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if (t.Left != nil) {
		Walk(t.Left, ch)
	}

	if (t.Right != nil) {
		Walk(t.Right, ch)
	}

	ch <- t.Value
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(tree.New(1), ch1)
	go Walk(tree.New(1), ch2)
	arr1, arr2 := []int{}, []int{}
	x := 0
	for i := 0; i < 20; i++ {
		select {
		case ch1 <- x:
			arr1 = append(arr1, x)
		case ch2 <- x:
			arr2 = append(arr2, x)
		}
	}

	function := func(i2, j int) bool {
		return i2 < j
	}
	sort.Slice(arr1, function)
	sort.Slice(arr2, function)

	for i := 0; i < 10; i++ {
		if (arr1[i] != arr2[i]) {
			return false
		}
	}

	return true
}

func main() {
	Same(tree.New(1), tree.New(1))
}
