/*
 * @lc app=leetcode.cn id=092 lang=golang
 *
 * [092] 反转链表 II
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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}

	// at least 2 nodes
	dummy := &ListNode{0, head}
	var preBound, postBound *ListNode
	ptr := dummy
	cursorIdx := 1
	for ; cursorIdx < m; cursorIdx++ {
		ptr = ptr.Next
	}
	preBound = ptr

	// reverse sub list
	subHead, subTail, postBound := reverseSub(ptr.Next, n - m +1)

	if subHead == nil || subTail == nil {
		// head为nil => m 是最后一个节点序号，则不翻转
		return head
	}

	preBound.Next = subHead
	subTail.Next = postBound
	return dummy.Next
}

// subHead 翻转后的head
// subTail 翻转后的tail
// postBound翻转后的tail后的节点
func reverseSub(head *ListNode, len int) (subHead *ListNode, subTail *ListNode, postBound *ListNode) {
	if head == nil {
		return head, head, head
	}

	dummy := &ListNode{0, head}
	pre, op, post := dummy, head, head.Next

	pre.Next = nil
	for i := 1; i <= len ; i++ {
		op.Next = pre
		pre = op
		op = post
		// 思路：根据“使用指针之前一定要保证不为空”规则
		// 正常的逻辑是循环跑len次，op指针从head一直跑到子串的尾巴节点，完成尾节点指针的转向，并且再向下移动一次
		// 但是post在这个过程中提前为nil了=>op操作到的子串位置的尾巴，并且这个尾节点也是总串的尾节点（才会出现post直接为空了）
		// 那么，正常流程继续执行的话，是pre移动到子串尾节点，op移动到postbound位置，post向后再移动个，所以，正常移动pre, op就行了
		// 指针规则2：（边界问题）根据其他指针规则，倒退边界场景出现的场景，以及应该进行的操作，保证边界
		// 指针问题思路：正向思考：从题目解答出发，想通思路，抽象方法；逆向思考：根据其他指针规则，倒退边界场景出现的场景，以及应该进行的操作，保证边界
		if post == nil {
			break
		}
		post = post.Next
	}

	head.Next = nil

	subTail = head
	subHead = pre
	postBound = op
	return subHead, subTail, postBound
}