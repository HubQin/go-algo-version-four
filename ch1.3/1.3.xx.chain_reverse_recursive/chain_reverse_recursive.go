package main

import "fmt"

// ListNode 定义单链表的节点结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 递归翻转单链表
func reverseList(head *ListNode) *ListNode {
	// 递归终止条件：当链表为空或只有一个节点时，直接返回该节点
	if head == nil || head.Next == nil {
		return head
	}

	// 递归翻转剩余的链表
	newHead := reverseList(head.Next)

	// 翻转当前节点与下一节点的指针
	head.Next.Next = head
	head.Next = nil

	// 返回新的头节点
	return newHead
}

// 打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// 创建链表: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("原链表:")
	printList(head)

	// 翻转链表
	reversedHead := reverseList(head)

	fmt.Println("翻转后的链表:")
	printList(reversedHead)
}
