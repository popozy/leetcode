/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 * 7月 09 2020
 */
package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)

	if leftHeight - rightHeight > 1 || rightHeight - leftHeight > 1 {
		return false
	} else {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}


// divide and conquer's template
// 递归返回条件
// 分段处理
// 合并结果
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMaxLength := maxDepth(root.Left)
	rightMaxLength := maxDepth(root.Right)

	if leftMaxLength > rightMaxLength {
		return leftMaxLength + 1
	} else {
		return rightMaxLength + 1
	}
}