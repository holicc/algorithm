package main

func main() {
	println(12 >> 1)
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
