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

	// fmt.Println(*linked_list.Retrieve(2).Val)

	// fmt.Println(linked_list.Update(0, 9))
	// fmt.Println(linked_list.Delete(1))
	// fmt.Println(linked_list.Size())

	// for _, element := range linked_list.FilterByVal(9) {
	// 	fmt.Println(*element.Val)
	// }

	// linked_list.Reverse()

	iter := linked_list.Head

	for iter != nil {
		fmt.Println(*iter.Val)

		iter = iter.Next
	}
}
