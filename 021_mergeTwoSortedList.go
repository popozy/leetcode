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

	// l1 l2 not empty both
	dummy := &ListNode{0, l1}
	pl1 := dummy
	pivot := pl1
	pl2 := l2
	for pl2 != nil {
		pl1 = pivot
		for pl1.Next != nil {
			if pl2.Val < pl1.Next.Val {
				pl1.Next = &ListNode{pl2.Val, pl1.Next}
				pivot = pl1
				break
			}
			pl1 = pl1.Next
		}
		// 跳出内层循环条件：1. pl1到尾部  2. 插入成功
		// 可知如果插入成功，则pl1.next不会为nil=> 1,2是互斥的=> pl1.next为nil时，当前的pl2还没插进去
		// pl2剩下的补到pl1尾巴即可
		if pl1.Next == nil {
			pl1.Next = pl2
			break
		}
		pl2 = pl2.Next
	}
	return dummy.Next
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	} else if l2 == nil {
// 		return l1
// 	}

// 	soldier := &ListNode{}
// 	mergedList := soldier

// 	for l1 != nil && l2 != nil {
// 		if l1.Val <= l2.Val {
// 			soldier.Val = l1.Val
// 			l1 = l1.Next
// 		} else {
// 			soldier.Val = l2.Val
// 			l2 = l2.Next
// 		}
// 		soldier.Next = &ListNode{}
// 		soldier = soldier.Next
// 	}

// 	if l1 == nil {
// 		soldier = getTail(mergedList)
// 		soldier.Next = l2
// 	} else if l2 == nil {
// 		soldier = getTail(mergedList)
// 		soldier.Next = l1
// 	}

// 	return mergedList
// }

// func getTail(listNode *ListNode) *ListNode {
// 	p := listNode
// 	if p.Next == nil {
// 		return nil
// 	}
// 	for ; p.Next.Next != nil; p = p.Next {}
// 	p.Next = nil
// 	return p
// }