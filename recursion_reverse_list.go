/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表-递归
 */
package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}
	rList := &ListNode{head.Val, nil}
	swapNode := head.Next
	head = head.Next

	return genRListByRecursion(rList, head, swapNode)
}

// 对于递归，确认以下几个要素
// 递归函数参数，一般是 1. 需要判定何时结束递归的flag 2. 其他需要在过程中变化使用的变量
// 递归函数处理好递归前的初始操作
// 递归函数本身，分两大分支： 1. 跳出递归 返回结果 2. 递归没结束时的基本操作 继续传参调用递归
func genRListByRecursion(rList *ListNode, head *ListNode, swapNode *ListNode) *ListNode {
	if head.Next == nil {
		// deal return
		head.Next = rList
		rList = head
		return rList
	} else {
		head = head.Next
		swapNode.Next = rList
		rList = swapNode
		swapNode = head
		return genRListByRecursion(rList, head, swapNode)
	}
}