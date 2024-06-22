package main

import (
	"fmt"
)

// ListNode 定义链表的结点结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除链表的第 k 个结点
func deleteNode(first *ListNode, k int) *ListNode {
	if first == nil {
		return nil // 链表为空，直接返回
	}

	if k <= 0 {
		return first // k 非法值，直接返回原链表
	}

	// 删除头结点的情况
	if k == 1 {
		return first.Next
	}

	current := first
	for i := 1; i < k-1; i++ {
		if current.Next == nil {
			return first // 如果 k 超出链表长度，返回原链表
		}
		// 移动current的指向， 直到第k-1个节点
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next // 删除第 k 个结点
	}

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

	// 删除第 k 个结点
	k := 3
	first = deleteNode(first, k)

	fmt.Println("删除第", k, "个结点后链表:")
	printList(first)
}
