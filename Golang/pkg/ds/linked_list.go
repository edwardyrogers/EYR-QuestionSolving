package ds

import (
	"eyr.question.solving/pkg/dt"
)

type LinkNode[T dt.Comparable] struct {
	Val  *T
	Next *LinkNode[T]
}

type LinkedList[T dt.Comparable] struct {
	Head *LinkNode[T]
}

func (std *LinkedList[T]) Insert(val T) {
	origin := &LinkNode[T]{}
	origin.Next = std.Head

	pointer := origin

	for pointer.Next != nil {
		pointer = pointer.Next
	}

	pointer.Next = &LinkNode[T]{&val, nil}

	std.Head = origin.Next
}

func (std *LinkedList[T]) Retrieve(val int) *LinkNode[T] {
	origin := &LinkNode[T]{}
	origin.Next = std.Head

	pointer := origin.Next

	pos := 0

	for pointer != nil {

		if pos == val {
			break
		}

		pointer = pointer.Next
		pos++
	}

	return pointer
}

func (std *LinkedList[T]) FilterByVal(val T) []LinkNode[T] {
	result := []LinkNode[T]{}

	origin := &LinkNode[T]{}
	origin.Next = std.Head

	pointer := origin.Next

	for pointer != nil {

		if *pointer.Val == val {
			result = append(result, *pointer)
		}

		pointer = pointer.Next
	}

	return result
}

func (std *LinkedList[T]) Update(idx int, newVal T) bool {
	result := false

	origin := &LinkNode[T]{}
	origin.Next = std.Head

	pointer := origin.Next

	pos := 0

	for pointer != nil {

		if pos == idx {
			pointer.Val = &newVal

			result = *pointer.Val == newVal

			break
		}

		pointer = pointer.Next
		pos++
	}

	std.Head = origin.Next

	return result
}

func (std *LinkedList[T]) Delete(idx int) bool {
	result := false

	origin := &LinkNode[T]{}
	origin.Next = std.Head

	previous := origin
	pointer := origin.Next

	pos := 0

	for pointer != nil {

		if pos == idx {
			previous.Next = pointer.Next
			pointer = previous.Next

			result = true
			break
		}

		previous = pointer
		pointer = pointer.Next
		pos++
	}

	std.Head = origin.Next

	return result
}

func (std *LinkedList[T]) Size() int {
	result := 0

	origin := &LinkNode[T]{}
	origin.Next = std.Head

	pointer := origin.Next

	for pointer != nil {
		pointer = pointer.Next
		result++
	}

	return result
}

func (std *LinkedList[T]) Reverse() {
	if std.Head == nil {
		return
	}

	var origin = &LinkNode[T]{}
	origin.Next = std.Head

	var leftEnd = origin

	var ptToStart = leftEnd.Next

	for ptToStart.Next != nil {
		var ptToCut = ptToStart.Next
		ptToStart.Next = ptToCut.Next
		ptToCut.Next = leftEnd.Next
		leftEnd.Next = ptToCut
	}

	std.Head = origin.Next
}
