package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 最长不重复子串
 */

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) > 0 && len(strings.Trim(s, " ")) == 0 {
		return 1
	} else if len(s) == 1 {
		return 1
	}

	var currentMaxLength int
	strArr := strings.Split(s, "")

	headIdx := 0
	tailIdx := 0
	subString := strArr[headIdx]

	for {
		tailIdx++
		if tailIdx >= len(strArr) - 1 {
			// deal last condition
			repeatIdx := strings.Index(subString, strArr[tailIdx])
			var lastSubLenghth int
			if repeatIdx == -1 {
				lastSubLenghth = tailIdx + 1 - headIdx
			} else {
				lastSubLenghth = tailIdx - headIdx
			}
			if lastSubLenghth > currentMaxLength {
				currentMaxLength = lastSubLenghth
			}
			break
		}

		repeatIdx := strings.Index(subString, strArr[tailIdx])
		if repeatIdx == -1 {
			// no repeat charactor
			subString = s[headIdx:tailIdx+1]
		} else {
			// charactor repeat
			if currentMaxLength < len(subString){
				currentMaxLength = len(subString)
			}
			headIdx = repeatIdx + headIdx + 1
			subString = s[headIdx:tailIdx+1]
		}
	}
	return currentMaxLength
}