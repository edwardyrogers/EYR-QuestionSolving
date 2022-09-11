package main

import (
	"fmt"

	"eyr.question.solving/pkg/ds"
	"eyr.question.solving/pkg/dt"
)

func main() {
	// https://leetcode.com/problems/add-two-numbers/

	listOne := ds.LinkedList[uint64]{}
	listOne.Insert(2)
	listOne.Insert(0)
	listOne.Insert(0)

	listTwo := ds.LinkedList[uint64]{}
	listTwo.Insert(3)
	listTwo.Insert(0)
	listTwo.Insert(0)

	iter := addTwoNumbers(listOne.Head, listTwo.Head)

	for iter != nil {
		fmt.Println(*iter.Val)

		iter = iter.Next
	}
}

func addTwoNumbers[T dt.Unsigned](l1 *ds.LinkedNode[T], l2 *ds.LinkedNode[T]) *ds.LinkedNode[T] {
	var carry T = 0
	var result = &ds.LinkedNode[T]{}
	var pointer = result

	for l1 != nil || l2 != nil || carry > 0 {

		var firstNum T = 0
		if l1 != nil {
			firstNum = *l1.Val
		}

		var secondNum T = 0
		if l2 != nil {
			secondNum = *l2.Val
		}

		sum := firstNum + secondNum + carry
		num := sum % 10
		carry = carry + (sum)/10

		pointer.Next = &ds.LinkedNode[T]{Val: &num, Next: nil}
		pointer = pointer.Next

		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}

	return result.Next
}
