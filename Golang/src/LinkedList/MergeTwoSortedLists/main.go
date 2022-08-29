package main

import (
	"fmt"
	"strconv"

	ds "eyr.question.solving/pkg/ds"
)

func main() {
	// https://leetcode.com/problems/merge-two-sorted-lists/

	listOne := ds.LinkedList{}
	listOne.Insert(1)
	listOne.Insert(3)
	listOne.Insert(5)

	listTwo := ds.LinkedList{}
	listTwo.Insert(2)
	listTwo.Insert(4)
	listTwo.Insert(6)

	iter := mergeTwoLists(listOne.Head, listTwo.Head)

	for iter != nil {
		fmt.Println(iter.Val)

		iter = iter.Next
	}
}

func mergeTwoLists(list1 *ds.LinkNode, list2 *ds.LinkNode) *ds.LinkNode {

	var result = &ds.LinkNode{Val: 0, Next: nil}
	var pointer = result

	for list1 != nil && list2 != nil {

		figure1Str := fmt.Sprint(list1.Val)
		figure1, _ := strconv.ParseUint(figure1Str, 0, 10)

		figure2Str := fmt.Sprint(list2.Val)
		figure2, _ := strconv.ParseUint(figure2Str, 0, 10)

		if figure1 < figure2 {
			pointer.Next = list1
			list1 = list1.Next
		} else {
			pointer.Next = list2
			list2 = list2.Next
		}

		pointer = pointer.Next
	}

	if list1 != nil {
		pointer.Next = list1
	} else if list2 != nil {
		pointer.Next = list2
	}

	return result.Next
}
