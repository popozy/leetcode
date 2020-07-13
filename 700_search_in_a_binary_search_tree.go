/*
 * @lc app=leetcode.cn id=700 lang=golang
 *
 * [700] 二叉搜索树中的搜索
 * 7月 13 2020
 */
package leetcode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil  {
		return nil
	}

	if root.Val == val {
		return root
	} else if val < root.Val {
		return searchBST(root.Left, val)
	} else if val > root.Val {
		return searchBST(root.Right, val)
	}
	return nil
}
