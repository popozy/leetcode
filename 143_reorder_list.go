/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] reorder-list
 * 7月 16 2020
 */
package leetcode

// 时间复杂度0(n)
// 实现：快慢指针找中点   翻转单向链表  简化版二路归并 边界处理
// 易出问题点：指针边界;分割链表时要置空;翻转单向链表时，要cut尾部调哑巴节点，处理好最后的节点
// golang的函数传参，是值传递
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// at least 2 nodes
	mid, newListTail := splitList(head)
	if newListTail != nil {
		newListTail.Next = nil
	}

	// reverse list
	mid = reverseList(mid)

	// merge  mid && head will not be nil
	head = mergeList(head, mid, newListTail)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// at least 2 nodes
	// make dummy at least 3 nodes
	dummy := &ListNode{0, head}
	pre, op, post := dummy, head, head.Next
	pre.Next = nil

	for post != nil {
		// change pointer to pre one
		op.Next = pre
		// move three nodes to next
		pre = op
		op = post
		post = post.Next
	}
	// break op meets last nodes
	// deal last node
	op.Next = pre
	head.Next = nil

	return op
}

// 二路归并
func mergeList(head *ListNode, mid *ListNode, tail *ListNode) *ListNode {
	// head mid will not be nil
	// deal
	if head.Next == nil {
		head.Next = mid
		mid.Next = tail
		return head
	}

	midPtr := mid
	headPtr := head
	// head and mid has at least 2 nodes respectively
	for midPtr != nil {
		// get isolate node in mid header
		tmp := midPtr
		// move midPtr to next
		midPtr = midPtr.Next

		// make midPtr 1st node point to headPtr.next
		tmp.Next = headPtr.Next
		// make head.next point to midPtr 1st node
		headPtr.Next = tmp

		headPtr = headPtr.Next.Next
	}

	ptr := head
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = tail
	return head
}

// return param1: 切分后需要倒排的len/2位置的首节点
// return param2: 当len为奇数时候，正中间那个节点，如果len为偶数，返回nil
func splitList(head *ListNode) (*ListNode, *ListNode) {
	// head  head.next check has been finished
	fast, slow := head, head

	// at least 2 nodes
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// odd nodes
	if fast.Next == nil {
		ptr := head
		for ptr.Next != slow {
			ptr = ptr.Next
		}
		ptr.Next = nil
		return slow.Next, slow
	}

	// even nodes
	newHead := slow.Next
	slow.Next = nil
	return newHead, nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reorderList(head *ListNode)  {
// 	if head == nil || head.Next == nil {
// 		return
// 	}
//
// 	length := 0
//
// 	for ptr := head; ptr != nil; {
// 		ptr = ptr.Next
// 		length++
// 	}
//
// 	newList := &ListNode{0, nil}
// 	pn := newList
// 	po := head
// 	halfLen := length /2
// 	for i := 1; i <= halfLen; i++ {
// 		peerNode := getPeerNode(head, length+1-i)
// 		// no need to check nil because it has done
// 		// append
// 		pn.Next = &ListNode{po.Val, &ListNode{peerNode.Val, nil}}
// 		pn = pn.Next.Next
// 		// no need to check because halfLength limitimg
// 		po = po.Next
// 	}
// 	if halfLen * 2 < length {
// 		pn.Next = &ListNode{po.Val, nil}
// 	}
// 	head = newList.Next
// 	for head != nil {
// 		print(head.Val)
// 		head = head.Next
// 	}
// }
//
// func getPeerNode(head *ListNode, targetIdx int) *ListNode {
// 	if head == nil {
// 		return head
// 	}
//
// 	// head.len >= targetIdx
// 	for i := 1; i < targetIdx; i++ {
// 		head = head.Next
// 	}
//
// 	return head
// }