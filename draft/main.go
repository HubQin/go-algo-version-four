package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Queue represents a queue that holds a slice of strings.
type Queue struct {
	items []string
}

// NewQueue creates a new Queue instance.
func NewQueue() *Queue {
	return &Queue{items: []string{}}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue) Enqueue(item string) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue.
func (q *Queue) Dequeue() string {
	if len(q.items) == 0 {
		return ""
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Size returns the number of items in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}

// Peek returns the item from the front of the queue without removing it.
func (q *Queue) Peek() string {
	if len(q.items) == 0 {
		return ""
	}
	return q.items[0]
}

func main() {
	// Parse the command line argument k.
	k := flag.Int("k", 0, "the position from the end of the queue")
	flag.Parse()

	if *k <= 0 {
		fmt.Println("Please provide a positive integer for k.")
		return
	}

	// Create a new queue.
	queue := NewQueue()

	// Read strings from standard input and enqueue them.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter strings (Press Ctrl+D to end input):")
	for scanner.Scan() {
		line := scanner.Text()
		queue.Enqueue(line)
	}

	// Check if the queue has at least k items.
	if queue.Size() < *k {
		fmt.Println("The input does not contain enough strings.")
		return
	}

	// Calculate the index of the k-th last element.
	index := queue.Size() - *k

	// Dequeue elements until we reach the desired element.
	for i := 0; i < index; i++ {
		queue.Dequeue()
	}

	// Print the k-th last element.
	fmt.Println("The k-th last string is:", queue.Peek())
}
