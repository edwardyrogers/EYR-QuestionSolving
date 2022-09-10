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

func reverseBetween[T dt.Comparable](head *ds.LinkNode[T], left int, right int) *ds.LinkNode[T] {
	if head == nil || left == right {
		return head
	}

	result := &ds.LinkNode[T]{}
	result.Next = head

	leftEnd := result

	pos := 1
	for pos < left {
		leftEnd = leftEnd.Next
		pos++
	}

	ptToStart := leftEnd.Next

	//       |     |
	// 1, 2, 3, 4, 5, 6, 7
	for left < right {

		ptToCut := ptToStart.Next
		ptToStart.Next = ptToCut.Next
		ptToCut.Next = leftEnd.Next
		leftEnd.Next = ptToCut

		left++
	}

	return result.Next
}
