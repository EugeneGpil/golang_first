package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	walkRecurse(t, ch)
	close(ch)
}

func walkRecurse(t *tree.Tree, ch chan int) {
	if (t.Left != nil) {
		walkRecurse(t.Left, ch)
	}

	if (t.Right != nil) {
		walkRecurse(t.Right, ch)
	}

	ch <- t.Value
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	arr1, arr2 := []int{}, []int{}
	for ch1Var := range ch1 {
		arr1 = append(arr1, ch1Var)
	}
	for ch2Var := range ch2 {
		arr2 = append(arr2, ch2Var)
	}
	sort.Ints(arr1)
	sort.Ints(arr2)

	for i := 0; i < 10; i++ {
		if (arr1[i] != arr2[i]) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
