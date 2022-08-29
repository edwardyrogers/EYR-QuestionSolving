package main

import (
	"fmt"

	ds "eyr.question.solving/pkg/ds"
)

func main() {
	// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

	list := ds.LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	iter := removeNthFromEnd(list.Head, 2)

	for iter != nil {
		fmt.Println(iter.Val)
		// fmt.Println("")

		iter = iter.Next
	}
}

func removeNthFromEnd(head *ds.LinkNode, n int) *ds.LinkNode {
	if n < 1 {
		return head
	}

	var result = &ds.LinkNode{Val: nil, Next: head}
	var quickNode = result
	var slowNode = result

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
