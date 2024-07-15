package main

import "fmt"

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

// 反转链表函数
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		tempNext := curr.Next // 暂存下一个节点
		curr.Next = prev      // 将当前节点的Next指针指向前一个节点
		prev = curr           // 前一个节点更新为当前节点
		curr = tempNext       // 当前节点更新为下一个节点
	}
	return prev
}

func main() {
	// 创建链表 1 -> 2 -> 3 -> 4 -> 5 -> nil
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}

	fmt.Println("Original List:")
	printList(head)

	// 反转链表
	reversedHead := reverseList(head)

	fmt.Println("Reversed List:")
	printList(reversedHead)
}
