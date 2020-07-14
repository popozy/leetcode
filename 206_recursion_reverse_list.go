/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表-递归
 */
package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	res, newHead := reverseRecursive(head.Next)
	res.Next = head
	head.Next = nil
	return newHead
}

// return param1: 1st node
// return param2: last node---new head
func reverseRecursive(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}

	res, newHead := reverseRecursive(head.Next)
	res.Next = head
	return head, newHead
}

// // 单向链表可以在指针头增加dummy节点，翻转链表时要先把dummy的next置空，防止后面cut的时候找不到尾巴
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	dummy := &ListNode{0, head}
// 	pre, op, post := dummy, head, head.Next
//
// 	// pre.Next = nil
// 	for post != nil {
// 		op.Next = pre
// 		pre = op
// 		op = post
// 		post = post.Next
// 	}
// 	// deal last node
// 	op.Next = pre
//
// 	// cut dummy node
// 	ptr := op
// 	for ptr.Next.Next != nil {
// 		ptr = ptr.Next
// 	}
// 	ptr.Next = nil
//
// 	return op
// }