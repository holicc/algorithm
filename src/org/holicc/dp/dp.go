package main

import "strconv"

//把数字翻译成字符串
//https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
//时间复杂度为 O(n)O(n)，空间复杂度为 O(n)O(n)
//动态规划，滚动数组
func translateNum(num int) int {
	s := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := range s {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := s[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}

//最长不含重复字符的子字符串
//https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
//时间复杂度 O(N) 其中 NN 为字符串长度，动态规划需遍历计算 dpdp 列表。
//空间复杂度 O(1) 字符的 ASCII 码范围为 00 ~ 127127 ，哈希表 dicdic 最多使用 O(128) = O(1)O(128)=O(1) 大小的额外空间。
func lengthOfLongestSubstring(s string) int {
	table := make(map[rune]int, 128)
	res, tmp := 0, 0
	for j, c := range s {
		i := -1
		if v, ok := table[c]; ok {
			i = v
		}
		table[c] = j
		//
		if tmp < j-i {
			tmp = tmp + 1
		} else {
			tmp = j - i
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

func pow3N(n int) int {
	o := 1
	for i := 0; i < n; i++ {
		o = (o * 3) % 1000000007
	}
	return o
}

//剪绳子 II
//https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	if m := 2 - n%3; m == 2 {
		return pow3N(n / 3)
	} else {
		return pow3N(n/3-m) * (m + 1) * 2 % 1000000007
	}
}
