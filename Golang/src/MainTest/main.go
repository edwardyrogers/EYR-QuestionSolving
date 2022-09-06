package main

import (
	"fmt"

	"eyr.question.solving/pkg/algo"
)

func main() {
	list := []string{"1", "2", "3", "4", "5"}
	// list := []int{1, 2, 3, 4, 5}

	fmt.Println(*algo.BinarySearch(list, "4"))
}
