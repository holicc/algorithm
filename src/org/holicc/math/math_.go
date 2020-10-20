package math

// https://leetcode-cn.com/problems/water-bottles/
// 换酒问题
func numWaterBottles(numBottles int, numExchange int) int {
	if numBottles >= numExchange {
		return (numBottles-numExchange)/(numExchange-1) + 1 + numBottles
	} else {
		return numBottles
	}
}

// https://leetcode-cn.com/problems/deep-dark-fraction/submissions/
// 分式化简
func fraction(cont []int) []int {
	l := len(cont)
	result := []int{cont[l-1], 1}
	for i := l - 2; i >= 0; i-- {
		result[0], result[1] = cont[i]*result[0]+result[1], result[0]
	}
	return result
}
