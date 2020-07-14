/*
 * @lc app=leetcode.cn id=082 lang=golang
 *
 * [082] leetcodeName
 * 7月 13 2020
 */
package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	// deleteWithoutNewList(head)
	// deleteWithNewList(head)
	return deleteWithDummy(head)
}

/** dummy节点可以用来处理边界条件 **/
func deleteWithDummy(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{0, nil}
	dummy.Next = head
	head = dummy
	pre, mid, post := head, head.Next, head.Next.Next
	var lastRm int

	for {
		if post == nil {
			// deal with last node: mid
			if mid.Val == lastRm {
				pre.Next = nil
			}
			break
		}
		if mid.Val == post.Val {
			lastRm = mid.Val
			for lastRm == mid.Val && post != nil {
				mid = post
				post = post.Next
			}
			pre.Next = mid
		} else {
			pre = mid
			mid = post
			post = post.Next
		}
	}

	return dummy.Next

}

// func deleteWithNewList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil{
// 		return head
// 	}
// 	newList := new(ListNode)
// 	newPtr := newList
// 	slow, fast := head, head.Next
//
// 	if fast.Next == nil {
// 		if slow.Val == fast.Val {
// 			return nil
// 		} else {
// 			return &ListNode{slow.Val, &ListNode{fast.Val, nil}}
// 		}
// 	}
//
// 	pre, slow ,fast := head, head.Next, head.Next.Next
// 	if pre.Val != slow.Val {
// 		newPtr.Next = &ListNode{pre.Val, nil}
// 		newPtr = newPtr.Next
// 	}
// 	for {
// 		if fast == nil {
// 			if slow.Val != pre.Val {
// 				newPtr.Next = &ListNode{slow.Val, nil}
// 			}
// 			break
// 		}
// 		if slow.Val != pre.Val && slow.Val != fast.Val {
// 			newPtr.Next = &ListNode{slow.Val, nil}
// 			newPtr = newPtr.Next
// 		}
// 		pre = slow
// 		slow = fast
// 		fast = fast.Next
// 	}
// 	return newList.Next
// }

// func deleteWithoutNewList(head *ListNode) *ListNode {
// 	for head == nil {
// 		return nil
// 	}
//
// 	slow, fast := head, head.Next
// 	repeatMap := make(map[int]struct{}, 0)
// 	for {
// 		if fast == nil {
// 			break
// 		}
// 		if slow.Val == fast.Val {
// 			slow.Next = fast.Next
// 			repeatMap[slow.Val] = struct{}{}
// 		} else {
// 			slow = fast
// 		}
// 		fast = fast.Next
// 	}
//
// 	// find the 1st valid node
// 	slow = head
// 	for {
// 		if slow == nil {
// 			break
// 		}
// 		if _, ok := repeatMap[slow.Val]; ok {
// 			slow = slow.Next
// 		} else {
// 			break
// 		}
// 	}
//
// 	head = slow
// 	// traverse the rest list node
// 	for {
// 		if slow == nil || slow.Next == nil {
// 			return head
// 		}
//
// 		if _, ok := repeatMap[slow.Next.Val]; ok {
// 			slow.Next = slow.Next.Next
// 		} else {
// 			// slow.next not in map
// 			slow = slow.Next
// 		}
// 	}
// }