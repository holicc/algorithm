package main

// https://leetcode-cn.com/problems/find-lucky-integer-in-an-array/
// 找出数组中的幸运数
func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	//
	var r = -1
	for k, v := range m {
		if k == v && r < k {
			r = k
		}
	}
	return r
}

// https://leetcode-cn.com/problems/special-array-with-x-elements-greater-than-or-equal-x/
// 特殊数组的特征值
func specialArray(nums []int) int {
	counting := make([]int, 1001)
	for i := range nums {
		counting[nums[i]]++
	}
	for i := len(counting) - 2; i > -1; i-- {
		counting[i] += counting[i+1]
	}
	for i := 0; i <= len(nums); i++ {
		if counting[i] == i {
			return i
		}
	}
	return -1
}

// https://leetcode-cn.com/problems/roman-to-integer/
// 罗马数字转整数
func romanToInt(s string) int {
	var _map = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// XXVII等于1+1+5+10+10 = 27 、IX等于10-1=9、XCI等于1+100-10。

	// 当右边的字符比左边的大， 说明是特殊组合
	pre, r := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := _map[s[i]]
		if cur >= pre {
			r += cur
		} else {
			r -= cur
		}
		pre = cur
	}
	return r
}
