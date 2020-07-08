/*
 * @lc app=leetcode.cn id=ID lang=golang
 *
 * [leetcodeID] basic operation--binary tree
 * 7月 08 2020
 */
package basic

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Value int
}

func main()  {
	tree := &TreeNode{nil, nil, 10}
	result := preOrderTraversal(tree)
	result2 := preOrderRecurse(tree)
	result3 := inOrderTraversal(tree)
	result4 := inOrderRecurse(tree)
	result5 := postOrderTraversal(tree)
	result6 := postOrderRecurse(tree)
	result7 := bfs(tree)
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)
}

func bfs(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	// queue为空证明全部都为空
	for len(queue) > 0 {
		valList := make([]int, 0)
		length := len(queue)
		// 遍历一层
		for i := 0; i < length; i++ {
			node := queue[0]
			valList = append(valList, node.Value)
			// pop first node in queue
			queue = queue[1:]

			// 把第二层塞到队列尾部
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, valList)
	}

	return result
}

func postOrderRecurse(treeNode *TreeNode) []int {
	result := make([]int, 0)
	result2 := make([]int, 0)
	result = postOrderDivideAndConquer(treeNode)
	postOrderDfs(treeNode, &result2)
	fmt.Println(result)
	fmt.Println(result2)
	return result
}

// DFS 深度搜索（从上到下） 和分治法区别：前者一般将最终结果通过指针参数传入，后者一般递归返回结果最后合并
func postOrderDfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	postOrderDfs(root.Left, result)
	postOrderDfs(root.Right, result)
	*result = append(*result, root.Value)
}

// 分治是先获取，再合并
func postOrderDivideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	result = append(result, divideAndConquer(root.Left)...)
	result = append(result, divideAndConquer(root.Right)...)
	result = append(result, root.Value)
	return  result
}

// 核心就是：根节点必须在右节点弹出之后，再弹出
func postOrderTraversal(treeNode *TreeNode) []int {
	if treeNode == nil {
		return nil
	}
	result := make([]int, 0)

	// 存储访问过但是不能pop到result的节点
	stack := make([]*TreeNode, 0)

	// 记录上次访问过的节点--用来判断root的right sub tree是否被访问过
	var lastVisit *TreeNode

	// stack不空或者当前存在访问节点都是要继续的
	for treeNode != nil || len(stack) != 0 {
		// 一路访问左子树下去，直到treeNode的最左侧子叶的空左子叶
		for treeNode != nil {
			stack  = append(stack, treeNode)
			treeNode = treeNode.Left
		}

		// 先取出来倒数第一个
		node := stack[len(stack)-1]

		// node右子叶为空 或者 右子叶访问过了 => 按照后续遍历规则，那就可以pop这个node了
		if node.Right == nil || node.Right == lastVisit {
			// 存结果
			result = append(result, node.Value)
			// pop
			stack = stack[:len(stack)-1]
			// 记录这个是上次被pop的
			// 设计思路：此时treeNode一定为nil; stack里当前node（非root）的前一个节点一定是node父节点
			// 所以，下一轮循环的node一定是当前node的父节点，那么需要标记lastVisit是node的右儿子来进入这个block
			lastVisit = node
		} else {
			// 如果node的右子叶不空 && 右子叶没有访问过 ==> 根据后续遍历规则，重置root节点，访问右子树
			treeNode = node.Right
		}
	}
	return result
}

func inOrderRecurse(treeNode *TreeNode) []int {
	if treeNode == nil {
		return nil
	}

	result := make([]int, 0)
	result = append(result, inOrderRecurse(treeNode.Left)...)
	result = append(result, treeNode.Value)
	result = append(result, inOrderRecurse(treeNode.Right)...)
	return result
}

func inOrderTraversal(treeNode *TreeNode) []int {
	if treeNode == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for treeNode != nil || len(stack) != 0 {
		for treeNode != nil {
			stack = append(stack, treeNode)
			treeNode = treeNode.Left
		}

		lastNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, lastNode.Value)
		treeNode = lastNode.Right
	}
	return result
}

func preOrderRecurse(treeNode *TreeNode) []int {
	if treeNode == nil {
		return nil
	}
	result := make([]int, 0)
	result = append(result, treeNode.Value)
	result = append(result, preOrderRecurse(treeNode.Left)...)
	result = append(result, preOrderRecurse(treeNode.Right)...)
	return result
}

func preOrderTraversal(treeNode *TreeNode) []int {
	if treeNode == nil {
		return nil
	}

	// save result to return
	result := make([]int, 0)

	// save data need to deal
	stack := make([]*TreeNode, 0)

	for treeNode != nil || len(stack) != 0 {
		for treeNode != nil {
			result = append(result, treeNode.Value)
			stack = append(stack, treeNode)
			treeNode = treeNode.Left
		}

		// pop last one
		lastNode := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		treeNode = lastNode.Right
	}

	return result
}