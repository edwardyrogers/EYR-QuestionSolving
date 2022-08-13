package main

import (
	"fmt"
	"golang/pkg/data_structures"
	"strconv"
)

func main() {
	// https://leetcode.com/problems/merge-k-sorted-lists/

	listOne := data_structures.LinkedList{}
	listOne.Insert(0)
	listOne.Insert(2)
	listOne.Insert(5)

	listTwo := data_structures.LinkedList{}
	listTwo.Insert(0)

	listThree := data_structures.LinkedList{}
	listThree.Insert(-1)

	lists := []*data_structures.LinkNode{listOne.Head, listTwo.Head, listThree.Head}

	iter := mergeKLists(lists)

	for iter != nil {
		fmt.Println(iter.Val)

		iter = iter.Next
	}
}

func mergeKLists(lists []*data_structures.LinkNode) *data_structures.LinkNode {

	var result = &data_structures.LinkNode{Val: 0, Next: nil}
	var tempResult = result

	for idx := 0; idx <= len(lists)-1; idx += 2 {

		var listA *data_structures.LinkNode = lists[idx]
		var listB *data_structures.LinkNode = nil

		if ((idx + 1) % len(lists)) != 0 {
			listB = lists[idx+1]
		}

		mergedLists := mergeTwoLists(listA, listB)
		tempResult.Next = mergeTwoLists(tempResult.Next, mergedLists)
	}

	return result.Next
}

func mergeTwoLists(list1 *data_structures.LinkNode, list2 *data_structures.LinkNode) *data_structures.LinkNode {

	var result = &data_structures.LinkNode{Val: 0, Next: nil}
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
