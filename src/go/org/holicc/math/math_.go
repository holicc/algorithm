package math

import (
	"fmt"
)

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

// https://leetcode-cn.com/problems/qi-wang-ge-shu-tong-ji/
// 期望个数统计
func expectNumber(scores []int) int {
	s := make(map[int]bool, len(scores))
	for _, i := range scores {
		s[i] = false
	}
	return len(s)
}

// https://leetcode-cn.com/problems/pascals-triangle-ii/
// 杨辉三角 II
func getRow(rowIndex int) []int {
	// 第0行
	nums := []int{1}
	for i := 1; i <= rowIndex; i++ {
		// 尾部追加1
		nums = append(nums, 1)
		// 倒序计算杨辉三角当前行
		for j := i - 1; j > 0; j-- {
			nums[j] += nums[j-1]
		}
	}
	return nums
}

func isHappy(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	}
	return n == 1
}

// https://leetcode-cn.com/problems/happy-number/
// 快乐数
func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// https://leetcode-cn.com/problems/day-of-the-week/
// 一周中的第几天
func dayOfTheWeek(day int, month int, year int) string {
	if month == 1 || month == 2 {
		month += 12
		year--
	}
	iWeek := (day + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400) % 7
	r := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	return r[iWeek]
}
