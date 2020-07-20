/*
 * @lc app=leetcode.cn id=094 lang=golang
 *
 * [094] binary-tree-inorder-traversal
 * 7月 20 2020
 */
package leetcode

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		lastNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, lastNode.Val)
		root = lastNode.Right
	}

	return result
}

// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}

// 	var result []int

// 	result = append(result, inorderTraversal(root.Left)...)
// 	result = append(result, root.Val)
// 	result = append(result, inorderTraversal(root.Right)...)
// 	return result
// }