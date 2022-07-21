package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(sl []int) *ListNode {
	var link *ListNode = nil
	for i := len(sl) - 1; i >= 0; i-- {
		node := ListNode{Val: sl[i], Next: link}
		link = &node
	}
	return link
}

func deleteDuplicates(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	link := node
	for {
		for node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
		}
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return link
}

func main() {
	list := createList([]int{0, 0, 0, 0, 0})

	list = deleteDuplicates(list)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
