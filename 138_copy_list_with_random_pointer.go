/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 * 7月 17 2020
 */
package leetcode

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	m := map[int]int{}
	for pK, idxK := head, 1; pK != nil; idxK++ {
		randomPointer := pK.Random
		// m[idxK]一定被更新，题目输入合法性保证
		if randomPointer == nil {
			// -1 meas nil
			m[idxK] = -1
		} else {
			for pV, idxV := head, 1; pV != nil; idxV++ {
				if randomPointer == pV {
					m[idxK] = idxV
					break
				}
				pV = pV.Next
			}
		}
		pK = pK.Next
	}

	// deep copy linked-list without updating randomPointer
	newList := new(Node)
	np := newList
	for ptr := head; ptr!= nil; ptr = ptr.Next {
		np.Next = &Node{ptr.Val, nil, nil}
		np = np.Next
	}
	newList = newList.Next

	// update random list o(n^2)
	for pk, idxK := newList, 1; pk != nil; idxK++ {
		idxV := m[idxK]
		if idxV == -1 {
			pk = pk.Next
			continue
		} else {
			// find target address by idxV
			pv := newList
			for i:= 1; i < idxV; i++ {
				pv = pv.Next
			}
			pk.Random = pv
		}
		pk = pk.Next
	}

	return newList
}