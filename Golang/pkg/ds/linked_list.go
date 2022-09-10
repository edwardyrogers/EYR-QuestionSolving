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
	if std.Head == nil {
		std.Head = &LinkNode[T]{&val, nil}

		return
	}

	if std.Head.Next == nil {
		std.Head.Next = &LinkNode[T]{&val, nil}

		return
	}

	next := std.Head.Next

	for next.Next != nil {
		next = next.Next
	}

	next.Next = &LinkNode[T]{&val, nil}
}

func (std *LinkedList[T]) Retrieve(val T) map[string]T {
	var resultSet map[string]T

	if std.Head != nil {
		index := 0

		if std.Head.Val == &val {
			resultSet = map[string]T{
				"idx": T(rune(index)),
				"val": *std.Head.Val,
			}
		}

		if resultSet == nil {
			next := std.Head.Next

			for next != nil {
				index++

				if next.Val == &val {
					resultSet = map[string]T{
						"idx": T(rune(index)),
						"val": *next.Val,
					}
				}

				next = next.Next
			}
		}
	}

	return resultSet
}

func (std *LinkedList[T]) FilterByVal(val T) []map[string]T {
	var resultSet []map[string]T

	if std.Head != nil {
		index := 0

		if std.Head.Val == &val {
			resultSet = append(resultSet, map[string]T{
				"idx": T(rune(index)),
				"val": *std.Head.Val,
			})
		}

		next := std.Head.Next

		for next != nil {
			index++

			if next.Val == &val {
				resultSet = append(resultSet, map[string]T{
					"idx": T(rune(index)),
					"val": *next.Val,
				})
			}

			next = next.Next
		}

	}

	return resultSet
}

func (std *LinkedList[T]) Update(oldVal T, newVal T) bool {
	result := false

	if std.Head != nil {

		if std.Head.Val == &oldVal {
			std.Head.Val = &newVal

			result = std.Head.Val == &newVal
		}

		if !result {
			next := std.Head.Next

			for next != nil {
				if next.Val == &oldVal {
					next.Val = &newVal

					result = next.Val == &newVal
					break
				}

				next = next.Next
			}
		}

	}

	return result
}

func (std *LinkedList[T]) Delete(val T) bool {
	result := false

	if std.Head != nil {

		if std.Head.Val == &val {
			if std.Head.Next != nil {
				std.Head = std.Head.Next
			} else {
				std.Head = nil
			}

			result = true
		}

		if !result {
			prev := std.Head
			next := std.Head.Next

			for next != nil {

				if next.Val == &val {
					if next.Next != nil {
						prev.Next = next.Next
					} else {
						prev.Next = nil
					}

					result = true
					break
				}

				prev = next
				next = next.Next
			}
		}
	}

	return result
}

func (std *LinkedList[T]) Size() int {
	var count []T

	if std.Head != nil {
		count = append(count, *std.Head.Val)

		next := std.Head.Next

		for next != nil {
			count = append(count, *std.Head.Val)
			next = next.Next
		}
	}

	return len(count)
}

func (std *LinkedList[T]) Reverse() {
	if std.Head == nil {
		return
	}

	var reaultHead = &LinkNode[T]{}
	reaultHead.Next = std.Head

	var leftEnd = reaultHead

	var ptToStart = leftEnd.Next

	for ptToStart.Next != nil {
		var ptToCut = ptToStart.Next
		ptToStart.Next = ptToCut.Next
		ptToCut.Next = leftEnd.Next
		leftEnd.Next = ptToCut
	}

	std.Head = reaultHead.Next
}
