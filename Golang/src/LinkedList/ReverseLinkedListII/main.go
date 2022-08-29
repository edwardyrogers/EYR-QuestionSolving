package main

import (
	"fmt"

	ds "eyr.question.solving/pkg/ds"
)

func main() {
	// https://leetcode.com/problems/reverse-linked-list-ii/

	list := ds.LinkedList{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	list.Insert(7)

	iter := reverseBetween(list.Head, 3, 5)

	for iter != nil {
		fmt.Println(iter.Val)

		iter = iter.Next
	}
}

func reverseBetween(head *ds.LinkNode, left int, right int) *ds.LinkNode {
	if head == nil || left == right {
		return head
	}

	var tempHead = &ds.LinkNode{}
	tempHead.Next = head

	var leftEnd = tempHead

	var pos = 1
	for pos < left {
		leftEnd = leftEnd.Next
		pos++
	}

	var ptToStart = leftEnd.Next

	//       |     |
	// 1, 2, 3, 4, 5, 6, 7
	for left < right {

		var ptToCut = ptToStart.Next
		ptToStart.Next = ptToCut.Next
		ptToCut.Next = leftEnd.Next
		leftEnd.Next = ptToCut

		left++
	}

	return tempHead.Next
}
