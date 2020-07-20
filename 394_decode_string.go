/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] decode-string
 * 7月 20 2020
 */
package leetcode

import (
	"strconv"
	"strings"
)

type Stack struct {
	stack []string
}

func decodeString(s string) string {
	// no need to validate because platform input guarantee
	stack := Constructor()
	arr := strings.Split(s, "")

	for _, v := range arr {
		if v != "]" {
			stack.Push(v)
		} else {
			// meet ] to close && pop char util [
			// get subStr
			subStr := ""
			for stack.Cap() > 0 &&stack.Top() != "[" {
				subStr = stack.Pop() + subStr
			}
			stack.Pop()

			// get factor
			factor := ""
			for stack.Cap() > 0 && stack.Top() >= "0" && stack.Top() <= "9" {
				factor = stack.Pop() + factor
			}
			factorInt, _ := strconv.Atoi(factor)
			result := ""
			for i := 1; i <= factorInt; i++ {
				result += subStr
			}
			// push new str back to arr
			stack.Push(result)
		}
	}
	result := ""
	for stack.Cap() > 0 {
		result = stack.Pop() + result
	}
	return result
}

func Constructor() Stack {
	return Stack{
		stack: []string{},
	}
}

func (this *Stack)Push(char string) {
	this.stack = append(this.stack, char)
}

func (this *Stack)Pop() string {
	char := this.stack[len(this.stack) - 1]
	this.stack = this.stack[:len(this.stack)-1]
	return char
}

func (this *Stack)Cap() int {
	return len(this.stack)
}

func (this *Stack)Top() string {
	return this.stack[len(this.stack) - 1]
}