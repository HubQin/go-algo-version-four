package main

import (
	"fmt"
)

// 使用递归方法获取链表中最大值

// ListNode 定义链表的结点结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// max 返回链表中键最大的结点的值 (递归方法)
func max(first *ListNode) int {
	if first == nil {
		return 0 // 链表为空，返回 0
	}

	// 递归获取剩余链表的最大值
	maxRest := max(first.Next)

	// 返回当前节点值与剩余链表最大值之间的较大者
	if first.Val > maxRest {
		return first.Val
	} else {
		return maxRest
	}
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

	// 获取链表中最大值
	maxVal := max(first)
	fmt.Println("链表中最大的结点值:", maxVal)
}

// 解:
// 对于链表中的每一个节点，最大值要么在当前节点，要么在剩余链表中。因此，问题可以分解为找当前节点值和剩余链表最大值中的较大值。
// 退出条件： 如果链表为空（即当前节点是 nil），返回 0。
// max(first) = MAX(first.Val, max(first.Next)
// =>max(first)  // first.Val = 1
// 		==> max(first.Next)  // first.Next.Val = 2
// 			===> max(first.Next.Next)  // first.Next.Next.Val = 3
// 				====> max(first.Next.Next.Next)  // first.Next.Next.Next.Val = 4
// 						=====> max(first.Next.Next.Next.Next)  // first.Next.Next.Next.Next = nil
// 						=====> returns 0
// 				====> returns max(4, 0) = 4
// 			===> returns max(3, 4) = 4
// 		==> returns max(2, 4) = 4
// => returns max(1, 4) = 4
