package main

import (
	"fmt"

	"eyr.question.solving/pkg/ds"
	"eyr.question.solving/pkg/dt"
)

func main() {
	// https://leetcode.com/problems/swap-nodes-in-pairs/

	list := ds.LinkedList[int]{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	list.Insert(7)

	iter := swapPairs(list.Head)

	fmt.Println()

	for iter != nil {
		fmt.Println(*iter.Val)

		iter = iter.Next
	}
}

func swapPairs[T dt.Comparable](head *ds.LinkedNode[T]) *ds.LinkedNode[T] {

	origin := &ds.LinkedNode[T]{}
	origin.Next = head

	pointer := origin

	pos := 0

	for pointer != nil {

		// 1234567
		ptToStart := pointer.Next
		checkNode := ptToStart != nil && ptToStart.Next != nil

		if pos%2 == 0 && checkNode {

			// 234567
			ptToCut := ptToStart.Next
			// 1 -> 34567
			ptToStart.Next = ptToCut.Next

			// 2 -> 1
			ptToCut.Next = pointer.Next
			// nil -> 2
			pointer.Next = ptToCut

		}

		pointer = pointer.Next
		pos++
	}

	return origin.Next
}
