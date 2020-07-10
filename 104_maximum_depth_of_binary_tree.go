/*
 * @lc app=leetcode.cn id=ID lang=golang
 *
 * [104] maxnmun_depth_of_binary_tree
 * 7月 08 2020
 */
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
