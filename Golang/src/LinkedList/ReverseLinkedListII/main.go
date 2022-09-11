package main

import (
	"fmt"

	"eyr.question.solving/pkg/ds"
	"eyr.question.solving/pkg/dt"
)

func main() {
	// https://leetcode.com/problems/reverse-linked-list-ii/

	list := ds.LinkedList[int]{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	list.Insert(7)

	iter := reverseBetween(list.Head, 3, 5)

	for iter != nil {
		fmt.Println(*iter.Val)

		iter = iter.Next
	}
}

func reverseBetween[T dt.Comparable](head *ds.LinkedNode[T], left int, right int) *ds.LinkedNode[T] {
	if head == nil || left == right {
		return head
	}

	result := &ds.LinkedNode[T]{}
	result.Next = head

	pointer := result

	pos := 1
	for pos < left {
		pointer = pointer.Next
		pos++
	}

	ptToStart := pointer.Next

	//       |     |
	// 1, 2, 3, 4, 5, 6, 7
	for left < right {

		// 234567
		ptToCut := ptToStart.Next
		// 1 -> 34567
		ptToStart.Next = ptToCut.Next

		// 2 -> 1
		ptToCut.Next = pointer.Next
		// nil -> 2
		pointer.Next = ptToCut

		left++
	}

	return result.Next
}
