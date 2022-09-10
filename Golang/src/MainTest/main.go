package main

import (
	"fmt"

	"eyr.question.solving/pkg/ds"
)

func main() {
	// list := []string{"1", "2", "3", "4", "5"}
	// list := []int{1, 2, 3, 4, 5}

	// fmt.Println(*algo.BinarySearch(list, "4"))

	linked_list := ds.LinkedList[int]{}
	linked_list.Insert(7)
	linked_list.Insert(8)
	linked_list.Insert(9)

	iter := linked_list.Head

	for iter != nil {
		fmt.Println(*iter.Val)

		iter = iter.Next
	}
}
