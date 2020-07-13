/*
 * @lc app=leetcode.cn id=002 lang=golang
 *
 * [002] 两数相加
 * 7月 13 2020
 */
package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultList := &ListNode{0, nil}
	returnList := resultList
	var sum int
	var adv int
	for l1 != nil || l2 != nil{
		factorA := getListVal(l1)
		factorB := getListVal(l2)
		sum, adv = addTwoFactors(factorA, factorB)
		resultList = setResult(resultList, sum, adv)
		if l1 != nil {
			l1 = l1.Next
		} else {
			factorA = 0
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			factorB = 0
		}
	}
	// deal highest node
	resultList = returnList
	for resultList.Next != nil {
		if resultList.Next.Val == 0 && resultList.Next.Next == nil {
			resultList.Next = nil
			break
		}
		resultList = resultList.Next
	}

	return returnList
}

// sum 当前位的值
// adv 需要进位的值
func setResult(listNode *ListNode, sum int, adv int) *ListNode {
	// 把当前位的值和低位进位的值相加，获取最终当前位的值，和加完需要向下一位进位的值
	nodeValue, nodeAdv := addTwoFactors(listNode.Val, sum)
	listNode.Val = nodeValue
	// nodeAdv 当前位相加需要进位的值
	// adv 需要进位的值
	listNode.Next = &ListNode{nodeAdv + adv, nil}
	return listNode.Next
}

// sum 当前位的值
// adv 需要进位的值
func addTwoFactors(factorA int, factorB int) (int, int) {
	sum := factorA + factorB
	adv := sum / 10
	sum = sum % 10
	return sum, adv
}

func getListVal(list *ListNode) int {
	if list != nil {
		factor := list.Val
		list = list.Next
		return factor
	}
	return 0
}