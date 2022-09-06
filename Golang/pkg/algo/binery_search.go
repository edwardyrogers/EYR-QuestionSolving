package algo

import (
	"eyr.question.solving/pkg/dt"
)

func BinarySearch[T dt.Comparable](input []T, target T) *T {
	left := 0
	right := len(input) - 1
	mid := 0

	for left < right {
		mid = left + (right-left)/2

		if target > input[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return &input[left]
}
