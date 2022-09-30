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

	for iter != nil {
		fmt.Println(*iter.Val)

		iter = iter.Next
	}
}

func swapPairs[T dt.Comparable](head *ds.LinkedNode[T]) *ds.LinkedNode[T] {

	origin := &ds.LinkedNode[T]{Val: nil}
	origin.Next = head

	pointer := origin

	pos := 1

	ptToConnect := pointer

	ptToStart := pointer.Next

	for ptToStart != nil && ptToStart.Next != nil {

		if pos == 2 {
			ptToConnect = ptToStart
			ptToStart = ptToStart.Next
			pos = 1

			continue
		}

		// 2...
		ptToCut := ptToStart.Next
		// 1 -> 3
		ptToStart.Next = ptToCut.Next

		// 2 -> 1
		ptToCut.Next = ptToConnect.Next
		// nil -> 2
		ptToConnect.Next = ptToCut

		pos++
	}

	return origin.Next
}
