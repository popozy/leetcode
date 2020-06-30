/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] mergeTwoSortedList
 * 6月 30 2020
 */
package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	soldier := &ListNode{}
	mergedList := soldier

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			soldier.Val = l1.Val
			l1 = l1.Next
		} else {
			soldier.Val = l2.Val
			l2 = l2.Next
		}
		soldier.Next = &ListNode{}
		soldier = soldier.Next
	}

	if l1 == nil {
		soldier = getTail(mergedList)
		soldier.Next = l2
	} else if l2 == nil {
		soldier = getTail(mergedList)
		soldier.Next = l1
	}

	return mergedList
}

func getTail(listNode *ListNode) *ListNode {
	p := listNode
	if p.Next == nil {
		return nil
	}
	for ; p.Next.Next != nil; p = p.Next {}
	p.Next = nil
	return p
}