package math

import "fmt"

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

// https://leetcode-cn.com/problems/fizz-buzz/
// Fizz Buzz
func fizzBuzz(n int) []string {
	list := make([]string, n)
	// 初始化赋值
	for i := 0; i < n; i, list[i] = i+1, fmt.Sprintf("%d", i+1) {
	}
	// 给3的倍数赋值
	for i := 2; i < n; i, list[i] = i+3, "Fizz" {
	}
	// 给5的倍数赋值
	for i := 4; i < n; i, list[i] = i+5, "Buzz" {
	}
	// 给3和5的倍数(即最小公倍数)赋值
	for i := 14; i < n; i, list[i] = i+15, "FizzBuzz" {
	}
	return list
}
