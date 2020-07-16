/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 判断回文链表
 * 7月 16 2020
 */
package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	reversedList := deepReverseList(head)
	ptrR := reversedList
	ptr := head
	for ptr != nil {
		if ptr.Val != ptrR.Val {
			return false
		}
		ptr = ptr.Next
		ptrR = ptrR.Next
	}
	return true
}

func deepReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newList := new(ListNode)
	ptrN := newList
	ptr := head
	for ptr != nil {
		ptrN.Next = &ListNode{ptr.Val, nil}
		ptr = ptr.Next
		ptrN = ptrN.Next
	}
	newList = newList.Next

	dummy := &ListNode{0, newList}
	pre, op, post := dummy, newList, newList.Next

	// reset pre.next
	pre.Next = nil
	for post != nil {
		// change op node pointer
		op.Next = pre
		// move 3 nodes to next block
		pre = op
		op = post
		post = post.Next
	}

	// break when op reach the last node
	op.Next = pre

	// cut dummy
	newList.Next = nil

	return op
}
