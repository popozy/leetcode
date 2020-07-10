/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 * 7月 09 2020
 */
package leetcode

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// divide and conquer's template
// 递归返回条件
// 分段处理
// 合并结果
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	leftSearch := lowestCommonAncestor(root.Left, p, q)
	rightSearch := lowestCommonAncestor(root.Right, p, q)

	if leftSearch != nil && rightSearch != nil {
		return root
	}

	if leftSearch != nil {
		return leftSearch
	}

	if rightSearch != nil {
		return rightSearch
	}

	return nil
}
