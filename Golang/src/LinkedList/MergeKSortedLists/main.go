package main

import (
	"fmt"

	"eyr.question.solving/pkg/ds"
	"eyr.question.solving/pkg/dt"
)

func main() {
	// https://leetcode.com/problems/merge-k-sorted-lists/

	listOne := ds.LinkedList[int]{}
	listOne.Insert(0)
	listOne.Insert(2)
	listOne.Insert(5)

	listTwo := ds.LinkedList[int]{}
	listTwo.Insert(0)

	listThree := ds.LinkedList[int]{}
	listThree.Insert(-1)

	lists := []*ds.LinkedNode[int]{listOne.Head, listTwo.Head, listThree.Head}

	iter := mergeKLists(lists)

	for iter != nil {
		fmt.Println(*iter.Val)

		iter = iter.Next
	}
}

func mergeKLists[T dt.Comparable](lists []*ds.LinkedNode[T]) *ds.LinkedNode[T] {

	result := &ds.LinkedNode[T]{}
	pointer := result

	for idx := 0; idx <= len(lists)-1; idx += 2 {

		var listA *ds.LinkedNode[T] = lists[idx]
		var listB *ds.LinkedNode[T] = nil

		if ((idx + 1) % len(lists)) != 0 {
			listB = lists[idx+1]
		}

		mergedLists := mergeTwoLists(listA, listB)
		pointer.Next = mergeTwoLists(pointer.Next, mergedLists)
	}

	return result.Next
}

func mergeTwoLists[T dt.Comparable](list1 *ds.LinkedNode[T], list2 *ds.LinkedNode[T]) *ds.LinkedNode[T] {

	result := &ds.LinkedNode[T]{}
	pointer := result

	for list1 != nil && list2 != nil {

		figure1 := *list1.Val
		figure2 := *list2.Val

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
