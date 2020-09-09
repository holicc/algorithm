package main

import (
	"math"
	"strconv"
)

func main() {
	kidsWithCandies([]int{2, 3, 1, 5, 3}, 3)
}

//数字序列中某一位的数字
//https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
func findNthDigit(n int) int {
	var digit = 1
	var count = 9
	var start = 1
	for n-count > 0 {
		n -= count
		digit += 1
		start *= 10
		count = digit * start * 9
	}
	var num = start + (n-1)/digit
	var res = strconv.Itoa(num)[(n-1)%digit] - '0'
	return int(res)
}

//https://leetcode-cn.com/problems/running-sum-of-1d-array/
//一维数组的动态和
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

//https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
//拥有最多糖果的孩子
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := math.MinInt32
	for i := range candies {
		if max < candies[i] {
			max = candies[i]
		}
	}
	//
	r := make([]bool, len(candies))
	for i := range r {
		r[i] = (candies[i] + extraCandies) >= max
	}
	return r
}
