package main

import (
	"fmt"
	"golang/pkg/data_structures"
)

func main() {
	// https://leetcode.com/problems/reverse-linked-list-ii/

	list := data_structures.LinkedList{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)

	result := reverseBetween(list.Head, 3, 5)

	fmt.Println(result.Val)
}

/**
 * Definition for singly-linked list.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */
func reverseBetween(head *data_structures.Node, left int, right int) *data_structures.Node {
	if head == nil {
		return nil
	}

	var curr *data_structures.Node = head
	var prev *data_structures.Node
	var next *data_structures.Node

	for curr.Next != nil {
		fmt.Println(curr.Val)
		fmt.Println(prev)
		fmt.Println(next)

		curr = curr.Next
	}

	return curr
}

func reverse() {
	// // Initialize current, previous and next pointers
	// Node* current = head;
	// Node* prev = NULL, *next = NULL;

	// while (current != NULL) {
	// 	// Store next
	// 	next = current->next;
	// 	// Reverse current node's pointer
	// 	current->next = prev;
	// 	// Move pointers one position ahead.
	// 	prev = current;
	// 	current = next;
	// }

	// head = prev;
}
