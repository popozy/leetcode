/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 检查环形链表进环的点
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
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	m := make(map[*ListNode]struct{}, 0)
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}