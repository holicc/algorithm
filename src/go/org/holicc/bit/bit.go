package main

import (
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	println(hammingWeight(0000000000000000000000000000101))
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

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/submissions/detail/111818144/
// 二进制链表转整数
func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result = (result << 1) | (head.Val)
		head = head.Next
	}
	return result
}

// https://leetcode-cn.com/problems/divisor-game/
// 除数博弈
func divisorGame(N int) bool {
	return N%2 == 0
}

// https://leetcode-cn.com/problems/number-of-1-bits/
// 位1的个数
func hammingWeight(num uint32) int {
	var r int
	for ; num != 0; r++ {
		num = num & (num - 1)
	}
	return r
}

// https://leetcode-cn.com/problems/add-digits/
// 各位相加 模九 数根
func addDigits(num int) int {
	return (num-1)%9 + 1
}

// https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/
// 根据数字二进制下 1 的数目排序
func sortByBits(arr []int) []int {
	ret := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ret[i] = bitCount(arr[i])*10000000 + arr[i]
	}
	sort.Ints(ret)
	for i := 0; i < len(arr); i++ {
		ret[i] = ret[i] % 10000000
	}
	return ret
}

// Java Integer.bitCount 实现
func bitCount(i int) int {
	i = i - ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f
	i = i + (i >> 8)
	i = i + (i >> 16)
	return i & 0x3f
}

// https://leetcode-cn.com/problems/reverse-bits/submissions/
// 颠倒二进制位
func reverseByte(b uint32, cache map[uint32]uint64) uint64 {
	value, ok := cache[b]
	if ok {
		return value
	}
	value = (uint64(b) * 0x0202020202 & 0x010884422010) % 1023
	cache[b] = value
	return value
}

func reverseBits(num uint32) uint32 {
	ret := uint64(0)
	power := uint64(24)
	var cache = map[uint32]uint64{}

	for num != 0 {
		ret += reverseByte(num&0xff, cache) << power
		num = num >> 8
		power -= 8
	}
	return uint32(ret)
}

// https://leetcode-cn.com/problems/binary-gap/
// 二进制间距
func binaryGap(N int) int {
	var last, ans int = -1, 0
	var i uint
	for i = 0; i < 32; i++ {
		if (N>>i)&1 > 0 {
			if last >= 0 {
				ans = max(ans, int(i)-last)
			}
			last = int(i)
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode-cn.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
// 将整数转换为两个无零整数的和
func getNoZeroIntegers(n int) []int {
	rand.Seed(time.Now().Unix())
	for {
		count := rand.Intn(n)
		count_string := strconv.Itoa(count)
		count2_string := strconv.Itoa(n - count)
		if !strings.Contains(count_string, "0") && !strings.Contains(count2_string, "0") {
			return []int{count, n - count}
		}
	}
}

// https://leetcode-cn.com/problems/binary-number-with-alternating-bits/submissions/
// 交替位二进制数
func hasAlternatingBits(n int) bool {
	n = n ^ (n >> 1)
	return (n & (n + 1)) == 0
}

// https://leetcode-cn.com/problems/complement-of-base-10-integer/
// 十进制整数的反码
func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}

	temp := 1
	for i := N; i > 0; i /= 2 {
		temp <<= 1
	}

	return (temp - 1) ^ N
}

// https://leetcode-cn.com/problems/sum-of-two-integers/solution/wei-yun-suan-by-ba-xiang/
// 两整数之和
func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		b = (a & b) << 1
		a = sum
	}
	return a
}

// https://leetcode-cn.com/problems/count-odd-numbers-in-an-interval-range/
// 在区间范围内统计奇数数目
func countOdds(low int, high int) int {
	/*
	   4种情况
	   l, r 分别为如下4种情况：
	   奇数    偶数    返回(h-l)>>1 + 1
	   奇数    奇数    返回(h-l)>>1 + 1
	   偶数    偶数    返回(h-l)>>1
	   偶数    奇数    返回(h-l)>>1 + 1
	*/
	if low&1 == 0 && high&1 == 0 {
		return (high - low) >> 1
	}
	return (high-low)>>1 + 1
}

// https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/solution/1689-shi-er-jin-zhi-shu-de-zui-shao-shu-1snv2/
// 十-二进制数的最少数目
func minPartitions(n string) int {
	max := -1
	for _, b := range n {
		if int(b-'0') > max {
			max = int(b - '0')
		}
	}
	return max
}
