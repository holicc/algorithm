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
