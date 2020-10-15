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
