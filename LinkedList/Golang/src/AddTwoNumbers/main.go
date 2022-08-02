package main

import (
	"fmt"
	"golang/pkg/data_structures"
)

func main() {
	// https://leetcode.com/problems/add-two-numbers/

	listOne := data_structures.LinkedList{}
	listOne.Insert(3)
	listOne.Insert(2)
	listOne.Insert(1)

	listTwo := data_structures.LinkedList{}
	listTwo.Insert(6)
	listTwo.Insert(5)
	listTwo.Insert(4)

	iter := addTwoNumbers(listOne.Head, listTwo.Head)

	for iter != nil {
		fmt.Println(iter.Val)

		iter = iter.Next
	}
}

func addTwoNumbers(l1 *data_structures.LinkNode, l2 *data_structures.LinkNode) *data_structures.LinkNode {

	return nil
}
