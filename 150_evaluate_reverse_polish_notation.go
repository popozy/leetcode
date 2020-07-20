/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] evaluate-reverse-polish-notation
 * 7月 20 2020
 */
package leetcode

import (
	"fmt"
	"strconv"
)

type Stack struct {
	stack []string
}

func evalRPN(tokens []string) int {
	// no need to validate tokens
	opList := "+-*/"
	factorStack := Constructor()

	for i := 0; i < len(tokens); i++ {
		char := tokens[i]
		if strings.Contains(opList, char) {
			// pop stack && calculate--update res
			// no need to do validate because input guarantee
			factorStack.calculate(char)
		} else {
			// push factor into stack
			factorStack.Push(char)
		}
	}
	result, _ := strconv.Atoi(factorStack.Pop())
	return result
}

func Constructor() Stack {
	stack := Stack{
		stack: []string{},
	}
	return stack
}

func (this *Stack)Pop() string {
	char := this.stack[len(this.stack) - 1]
	this.stack = this.stack[:len(this.stack)-1]
	return char
}

func (this *Stack)Top() string {
	return this.stack[len(this.stack) - 1]
}

func (this *Stack)Push(x string) {
	this.stack = append(this.stack, x)
}

// return new result
func (this *Stack) calculate(char string) {
	factor1, _ := strconv.Atoi(this.Pop())
	factor2, _:= strconv.Atoi(this.Pop())
	var newFactor int
	if char == "+" {
		newFactor = factor1 + factor2
	} else if char == "-" {
		newFactor = factor2 - factor1
	} else if char == "*" {
		newFactor = factor2 * factor1
	} else if char == "/" {
		newFactor = factor2 / factor1
	} else {
		fmt.Println("invalid operator")
	}
	this.Push(strconv.Itoa(newFactor))
}