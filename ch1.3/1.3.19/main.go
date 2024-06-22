package main

import (
	"fmt"
)

// ListNode 定义链表的结点结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除链表的尾结点
func deleteTailNode(first *ListNode) *ListNode {
	if first == nil {
		return nil // 链表为空，直接返回
	}

	if first.Next == nil {
		return nil // 链表只有一个结点，删除后返回nil
	}

	// 遍历到链表的倒数第二个结点
	current := first
	for current.Next.Next != nil {
		current = current.Next
	}

	// 删除尾结点
	current.Next = nil
	// 因为current和first都是指向链表节点的指针，所以通过current修改了最后一个节点， 也就是first指向的那个链表
	return first
}

// 打印链表
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// 创建链表 1 -> 2 -> 3 -> 4
	node4 := &ListNode{Val: 4, Next: nil}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	first := &ListNode{Val: 1, Next: node2}

	fmt.Println("原链表:")
	printList(first)

	// 删除尾结点
	first = deleteTailNode(first)

	fmt.Println("删除尾结点后链表:")
	printList(first)
}
