/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 * 7月 10 2020
 */
package leetcode

import "math"

// 思路：
//平衡二叉树的构建是拿着新node从root一路遍历下去，保证右子树最小值大于root，左子树最大值小于root
//因此，验证时不能只判断父子节点的大小关系，要以左右子树的最大最小值去和root比较，同时需要自上而下保证每个节点都符合前述规则--双层递归
// 递归：
//  触底条件：达到最底层叶子节点，返回true，因为叶子节点不用再和其子树比较，返回true不影响回溯的结果
//  分治：考虑单子树场景，获取最小/大值
//  合并（治理）：root和左右子树（如果有）的最大/小值比较，如果不符合条件，直接return false，符合特征，则递归获取左右子树是否符合bst
// 如果当前root判定、左子树、右子树均合法，则返回true
func isValidBST(root *TreeNode) bool {
	currentNodeValid := false
	leftMax, rightMin := math.MinInt64, math.MaxInt64

	if root == nil ||  root.Left == nil && root.Right == nil {
		return true
	} else if root.Left == nil {
		rightMin = minValueLeaf(root.Right)
	} else if  root.Right == nil {
		leftMax = maxValueLeaf(root.Left)
	} else {
		leftMax = maxValueLeaf(root.Left)
		rightMin = minValueLeaf(root.Right)
	}

	if root.Val > leftMax && root.Val < rightMin {
		currentNodeValid = true
	} else {
		return false
	}

	leftTreeValid := isValidBST(root.Left)
	rightTreeValid := isValidBST(root.Right)

	return leftTreeValid && rightTreeValid && currentNodeValid
}

func minValueLeaf(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	} else if root.Left == nil {
		rightMin := minValueLeaf(root.Right)
		if root.Val > rightMin {
			return rightMin
		}
		return root.Val
	} else if root.Right == nil {
		leftMin := minValueLeaf(root.Left)
		if root.Val > leftMin {
			return leftMin
		}
		return root.Val
	} else {
		leftMin := minValueLeaf(root.Left)
		rightMin := minValueLeaf(root.Right)
		min := root.Val
		if min > leftMin {
			min = leftMin
		}
		if min > rightMin {
			min = rightMin
		}
		return min
	}
}

// 查找到以root为根节点的bt的最大节点值
func maxValueLeaf(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return  root.Val
	} else if root.Left == nil {
		maxLeaf := maxValueLeaf(root.Right)
		if root.Val > maxLeaf {
			return root.Val
		}
		return maxLeaf
	} else if root.Right == nil {
		maxLeaf := maxValueLeaf(root.Left)
		if root.Val > maxLeaf {
			return root.Val
		}
		return maxLeaf
	} else {
		leftMax := maxValueLeaf(root.Left)
		rightMax := maxValueLeaf(root.Right)
		defaultMax := root.Val
		if defaultMax < leftMax {
			defaultMax = leftMax
		}
		if defaultMax < rightMax {
			defaultMax = rightMax
		}
		return defaultMax
	}
}
