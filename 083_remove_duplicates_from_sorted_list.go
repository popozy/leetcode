/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [083] 删除排序链表中的重复元素
 * 7月 13 2020
 */
package leetcode


// 单层循环
// 快慢指针用来解决链表操作中list的节点间需要比较的场景
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head.Next
	for {
		if fast == nil {
			return head
		}
		if slow.Val == fast.Val {
			slow.Next = fast.Next
		} else {
			slow = slow.Next
		}
		fast = fast.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 朴素解法
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil{
// 		return head
// 	}
//
// 	ptr := head
// 	for {
// 		for {
// 			if ptr == nil || ptr.Next == nil {
// 				return head
// 			}
// 			if ptr.Val == ptr.Next.Val {
// 				// remove next node
// 				ptr.Next = ptr.Next.Next
// 			} else {
// 				break
// 			}
// 		}
// 		ptr = ptr.Next
// 	}
// }