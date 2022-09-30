package main

import (
	"fmt"

	"eyr.question.solving/pkg/ds"
	"eyr.question.solving/pkg/dt"
)

func main() {
	// https://leetcode.com/problems/reverse-nodes-in-k-group/

	list := ds.LinkedList[int]{}

	list.Insert(1)
	list.Insert(2)
	// list.Insert(3)
	// list.Insert(4)
	// list.Insert(5)
	// list.Insert(6)
	// list.Insert(7)
	// list.Insert(8)

	iter := reverseKGroup(list.Head, 2)

	fmt.Println()

	for iter != nil {
		fmt.Println(*iter.Val)

		iter = iter.Next
	}
}

func reverseKGroup[T dt.Comparable](head *ds.LinkedNode[T], k int) *ds.LinkedNode[T] {

	origin := &ds.LinkedNode[T]{Val: nil}
	origin.Next = head

	pointer := origin

	ptToRunFirst := pointer

	ptToConnect := pointer

	ptToStart := pointer.Next

	pos := 1

	countToStop := 0

	for idx := 0; idx <= k; idx++ {
		ptToRunFirst = ptToRunFirst.Next
	}

	for ptToStart != nil && ptToStart.Next != nil {

		if ptToRunFirst != nil && ptToRunFirst.Next != nil {
			ptToRunFirst = ptToRunFirst.Next
		} else {
			countToStop++
		}

		if pos == k {

			ptToConnect = ptToStart
			ptToStart = ptToStart.Next

			if countToStop > 1 {
				break
			}

			pos = 1
			countToStop = 0

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
