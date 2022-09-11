package main

import (
	"fmt"

	"eyr.question.solving/pkg/ds"
	"eyr.question.solving/pkg/dt"
)

func main() {
	// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

	list := ds.LinkedList[int]{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	iter := removeNthFromEnd(list.Head, 2)

	for iter != nil {
		fmt.Println(*iter.Val)

		iter = iter.Next
	}
}

func removeNthFromEnd[T dt.Comparable](head *ds.LinkedNode[T], n int) *ds.LinkedNode[T] {
	if n < 1 {
		return head
	}

	result := &ds.LinkedNode[T]{Val: nil, Next: head}

	quickNode := result
	slowNode := result

	for idx := 1; idx <= n; idx++ {
		quickNode = quickNode.Next
	}

	for quickNode.Next != nil {
		quickNode = quickNode.Next
		slowNode = slowNode.Next
	}

	slowNode.Next = slowNode.Next.Next

	return result.Next
}
