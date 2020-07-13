/*
 * @lc app=leetcode.cn id=029 lang=golang
 *
 * [029] 两数相除
 * 7月 13 2020
 */
package leetcode

func divide(dividend int, divisor int) int {
	result := 0
	if dividend == 0 {
		return result
	}

	// dividend && divisor != 0
	isPositive := getPostiveRes(dividend, divisor)

	divisorPos, dividendPos := convert2Pos(divisor, dividend)

	if divisorPos > dividendPos {
		return 0
	}

	power := uint32(31)
	// get power: divisorPos * power(2, power) < dividendPos <= divisor * power(2, power+1)
	for {
		dividendMove := dividendPos >> power
		if dividendMove < divisorPos {
			power--
			continue
		} else {
			break
		}
	}

	base := calBase(divisorPos, power)

	modCount := calModCount(divisorPos, dividendPos, base)

	result = calPositiveRes(int(power), modCount)

	return dealResult(isPositive, result)
}

func convert2Pos(divisor int, dividend int) (int, int) {
	dividendPos := dividend
	divisorPos := divisor

	if dividend < 0 {
		dividendPos = 0 - dividend
	}
	if divisor < 0 {
		divisorPos = 0 - divisor
	}
	return divisorPos, dividendPos
}

func dealResult(isPositive bool, result int) int {
	MaxLimit := int(math.Pow(2, 31)) - 1
	MinLimit := 0 - int(math.Pow(2, 31))
	if result > MaxLimit {
		if isPositive {
			return MaxLimit
		} else {
			return MinLimit
		}
	}

	if isPositive {
		return result
	} else {
		return -result
	}
}

func calPositiveRes(power int, modCount int) int {
	// result = math.power(2, power) + modCount - 1
	result := 1
	for idx := 0; idx < power; idx++ {
		result = result << 1
	}

	result += modCount
	return result
}

func calModCount(divisorPos int, dividendPos int, base int) int {
	var modCount int
	for {
		base += divisorPos
		if base < dividendPos {
			modCount++
			continue
		} else if base == dividendPos {
			modCount++
			break
		}else{
			break
		}
	}
	return modCount
}

func calBase(divisorPos int, power uint32) int {
	// baseline: power dividendMove divisor dividendPos
	// break loop when divisorPos * power(2, power) + modCount * divisorPos > dividendPos
	base := divisorPos << 1
	powerIdx := int(power)
	powerIdx--
	for ; powerIdx > 0; powerIdx-- {
		base = base << 1
	}
	return base
}

func getPostiveRes(dividend int, divisor int) bool {
	isPositive := true
	if (dividend ^ divisor) < 0 {
		isPositive = false
	}
	return isPositive
}