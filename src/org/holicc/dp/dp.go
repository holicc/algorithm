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
