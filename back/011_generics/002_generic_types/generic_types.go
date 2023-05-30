package main

import (
	"fmt"
)

type MyList[T any] struct {
	next *MyList[T]
	val T
}

func (myListPointer *MyList[T]) getAllVals() []T {
	pointer := myListPointer

	res := make([]T, 0)

	for ; pointer != nil; pointer = pointer.next {
		res = append(res, pointer.val)
	}

	return res
}

func main() {
	myList := MyList[int] {
		val: 0,
		next: &MyList[int] {
			val: 1,
			next: &MyList[int] {
				val: 2,
			},
		},
	}

	fmt.Println(myList.getAllVals())
}
