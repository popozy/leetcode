/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] min-stack
 * 7月 20 2020
 */
package leetcode

type MinStack struct {
	stack []int
	minIdx int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minIdx: -1,
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if this.minIdx == -1 {
		this.minIdx = 0
	} else {
		if this.stack[this.minIdx] > x {
			this.minIdx = len(this.stack) - 1
		}
	}
}


func (this *MinStack) Pop()  {
	if this.minIdx == len(this.stack) - 1 {
		newMinIdx := 0
		for i := 0; i < this.minIdx; i++ {
			if this.stack[newMinIdx] > this.stack[i] {
				newMinIdx = i
			}
		}
		this.minIdx = newMinIdx
	}
	this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.stack[this.minIdx]
}