/*
 * @lc app=leetcode.cn id=005 lang=golang
 *
 * [005] 最长回文子串
 * 7月 13 2020
 */
package leetcode

import "strings"

func longestPalindrome(s string) string {
	strArr := strings.Split(s, "")
	var m []string
	if len(s) == 1 {
		return s
	} else if len(s) == 2 && strArr[0] != strArr[1] {
		return strArr[0]
	}

	for index, c := range strArr {
		// deal bound
		if index == 0 && strArr[index] == strArr[index + 1] ||
			index == len(strArr) -1 && strArr[index] == strArr[index -1] {
			m =append(m, c+c)
			continue
		} else if index == 0 || index == len(strArr) - 1 {
			continue
		}

		m = append(m, strachEvenPalin(index, strArr))
		m = append(m, strachOddPalin(index, strArr))
	}
	maxPalindrome := genMaxPalindrome(m)
	return maxPalindrome
}

func genMaxPalindrome(m []string) string {
	maxString := ""
	for _, item := range m {
		if len(item) > len(maxString) {
			maxString = item
		}
	}
	return maxString
}

// ... a a ...
func strachEvenPalin(index int, strArr []string) string {
	subStr := ""
	halfDur := 1
	for ; index + 1 - halfDur >= 0 && index + halfDur <= len(strArr) - 1 &&
		strArr[index + 1 - halfDur] == strArr[index + halfDur]; halfDur++ {
		subStr = strArr[index + 1 - halfDur] + subStr + strArr[index + halfDur]
	}
	return subStr
}

// ... a x a ...
func strachOddPalin(index int, strArr []string) string {
	subStr := strArr[index]
	halfDur := 1
	for ; index - halfDur >= 0 && index + halfDur <= len(strArr) - 1 &&
		strArr[index - halfDur] == strArr[index + halfDur]; halfDur++ {
		subStr = strArr[index-halfDur] + subStr + strArr[index+halfDur]
	}
	return subStr
}