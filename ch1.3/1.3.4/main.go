package main

import (
	"bufio"
	"fmt"
	"os"
)

// Stack structure
type Stack struct {
	elements []rune
}

// Push an element onto the stack
func (s *Stack) Push(r rune) {
	s.elements = append(s.elements, r)
}

// Pop an element from the stack
func (s *Stack) Pop() (rune, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	r := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return r, true
}

// Check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Check if the parentheses are balanced
func isBalanced(input string) bool {
	stack := &Stack{}

	for _, char := range input {
		switch char {
		// 遇到左边括号，则将其压入栈中
		case '(', '{', '[':
			stack.Push(char)
		case ')', '}', ']':
			if stack.IsEmpty() {
				return false
			}
			// 遇到右边括号，则从栈中弹出一个元素，并检查是否匹配
			top, _ := stack.Pop()
			if !isMatchingPair(top, char) {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

// Check if the given pair of parentheses is matching
func isMatchingPair(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '{':
		return close == '}'
	case '[':
		return close == ']'
	default:
		return false
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the text to check for balanced parentheses:")
	input, _ := reader.ReadString('\n')

	if isBalanced(input) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
