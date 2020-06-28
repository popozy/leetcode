/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表-迭代
 */

package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	rList := &ListNode{head.Val, nil}
	swapNode := head.Next
	head = head.Next

	for {
		if head.Next == nil {
			//deal the last node
			head.Next = rList
			rList = head
			return rList
		}

		head = head.Next
		swapNode.Next = rList

		//move pointer
		rList = swapNode
		swapNode = head
	}
}