/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
 * 7月 15 2020
 */
package leetcode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}