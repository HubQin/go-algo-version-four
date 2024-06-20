package main

import (
	"fmt"
)

func isValidSequence(sequence []int) bool {
	stack := []int{}
	n := 10 // Numbers are 0 to 9
	index := 0

	for i := 0; i < n; i++ {
		stack = append(stack, i) // Push operation

		// Pop from stack if the top of the stack matches the current sequence element
		for len(stack) > 0 && stack[len(stack)-1] == sequence[index] {
			stack = stack[:len(stack)-1] // Pop operation
			index++
		}
	}

	return len(stack) == 0 && index == len(sequence)
}

func main() {
	sequences := [][]int{
		{4, 3, 2, 1, 0, 9, 8, 7, 6, 5},
		{4, 6, 8, 7, 5, 3, 2, 9, 0, 1},
		{2, 5, 6, 7, 4, 8, 9, 3, 1, 0},
		{4, 3, 2, 1, 0, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 9, 8, 7, 0},
		{0, 4, 6, 5, 3, 8, 1, 7, 2, 9},
		{1, 4, 7, 9, 8, 6, 5, 3, 0, 2},
		{2, 1, 4, 3, 6, 5, 8, 7, 9, 0},
	}

	for i, sequence := range sequences {
		if isValidSequence(sequence) {
			fmt.Printf("Sequence %d is possible.\n", i+1)
		} else {
			fmt.Printf("Sequence %d is not possible.\n", i+1)
		}
	}
}
