package main

import (
	"math"
	"strings"
)

func main() {
	println(strToInt("-2147483647"))
}

//1～n整数中1出现的次数
//输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
//例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。
//https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/
func countDigitOne(n int) int {
	var ans int
	high, cur, low, digit := n/10, n%10, 0, 1
	for high != 0 || cur != 0 {
		if cur == 0 {
			ans += high * digit
		} else if cur == 1 {
			ans += high*digit + low + 1
		} else {
			ans += (high + 1) * digit
		}
		low = cur*digit + low
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return ans
}

//数值的整数次方
//https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
//时间复杂度 O(log2n)
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	//
	if n < 0 {
		x = 1 / x
		n = -n
	}
	//
	r := 1.0
	for n > 0 {
		if (n & 1) == 1 {
			r *= x
		}
		x *= x
		n >>= 1
	}
	return r
}

//把字符串转换成整数
//https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/
func strToInt(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	var sign = 1
	var i int
	var boundary = math.MaxInt32 / 10
	var res int
	// 如果是符号位则i++
	if str[i] == '-' || str[i] == '+' {
		if str[i] == '-' {
			sign = -1
		}
		i++
	}
	// 开始处理数值
	for ; i < len(str); i++ {
		// 不符合规则则跳过
		if str[i] < '0' || str[i] > '9' {
			break
		}
		// 处理越界的情况
		if res > boundary || (res == boundary && str[i] > '7') {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		res = res*10 + int(str[i]-'0')

	}
	return sign * res
}

//https://leetcode-cn.com/problems/xor-operation-in-an-array/
//数组异或操作
func xorOperation(n int, start int) int {
	var result int
	for i := 0; i < n; i++ {
		result ^= start + 2*i
	}
	return result
}