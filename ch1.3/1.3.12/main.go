package main

import "fmt"

// 编写一个可迭代的 Stack 用例，它含有一个静态的 copy() 方法，接受一个字符串的栈作为参数
// 并返回该栈的一个副本。注意：这种能力是迭代器价值的一个重要体现，因为有了它我们无需
// 改变基本 API 就能够实现这种功能。

// Stack represents a stack that holds a slice of strings.
type Stack struct {
	items []string
}

// NewStack creates a new Stack instance.
func NewStack() *Stack {
	return &Stack{items: []string{}}
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the item from the top of the stack without removing it.
func (s *Stack) Peek() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

// Iterator returns a channel that can be used to iterate over the stack.
func (s *Stack) Iterator() <-chan string {
	ch := make(chan string)
	go func() {
		for i := len(s.items) - 1; i >= 0; i-- {
			ch <- s.items[i]
		}
		close(ch)
	}()
	return ch
}

// Copy creates a copy of the given stack.
func Copy(s *Stack) *Stack {
	newStack := NewStack()
	for item := range s.Iterator() {
		newStack.Push(item)
	}
	// Since items are pushed in reverse order during iteration, we need to reverse them back.
	reversedItems := []string{}
	for item := range newStack.Iterator() {
		reversedItems = append([]string{item}, reversedItems...)
	}
	newStack.items = reversedItems
	return newStack
}

func main() {
	// Example usage
	stack := NewStack()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	fmt.Println("Original Stack:")
	for item := range stack.Iterator() {
		fmt.Println(item)
	}

	copiedStack := Copy(stack)

	fmt.Println("Copied Stack:")
	for item := range copiedStack.Iterator() {
		fmt.Println(item)
	}
}
