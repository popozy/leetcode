/*
 * @lc app=leetcode.cn id=086 lang=golang
 *
 * [086] 分隔链表
 * 7月 15 2020
 */
package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	pivot := &ListNode{x, nil}
	dummy := &ListNode{0, pivot}
	pre := dummy

	ptr := head
	for ptr != nil {
		if ptr.Val < x {
			tmp := &ListNode{ptr.Val, pivot}
			pre.Next = tmp
			pre = tmp
		} else {
			tmp := pivot
			for tmp.Next != nil {
				tmp = tmp.Next
			}
			tmp.Next = &ListNode{ptr.Val, nil}
		}
		ptr = ptr.Next
	}
	// remove pivot
	pre.Next = pre.Next.Next

	return dummy.Next
}