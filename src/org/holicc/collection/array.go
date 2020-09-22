package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
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

//https://leetcode-cn.com/problems/number-of-good-pairs/
//好数对的数目
func numIdenticalPairs(nums []int) int {
	var r int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				r++
			}
		}
	}
	return r
}

//https://leetcode-cn.com/problems/shuffle-the-array/
//重新排列数组
//给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。
//请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组。
func shuffle(nums []int, n int) []int {
	r := make([]int, len(nums))
	for j, k, i := n, 0, 0; i < len(r); i++ {
		if (i & 1) == 1 {
			r[i] = nums[j]
			j++
		} else {
			r[i] = nums[k]
			k++
		}
	}
	return r
}

//https://leetcode-cn.com/problems/guess-numbers/submissions/
//猜数字
func game(guess []int, answer []int) int {
	r := 0
	for i := range answer {
		if guess[i] == answer[i] {
			r++
		}
	}
	return r
}

//https://leetcode-cn.com/problems/na-ying-bi/
//桌上有 n 堆力扣币，每堆的数量保存在数组 coins 中。我们每次可以选择任意一堆，拿走其中的一枚或者两枚，求拿完所有力扣币的最少次数。
func minCount(coins []int) int {
	r := 0
	for i := range coins {
		r += coins[i]/2 + coins[i]%2
	}
	return r
}

// https://leetcode-cn.com/problems/matrix-diagonal-sum/
// 矩阵对角线元素的和
func diagonalSum(mat [][]int) int {
	var result int
	for i, val := range mat {
		if i == len(mat)-i-1 {
			result += val[i]
		} else {
			result += val[i] + val[len(mat)-i-1]
		}
	}
	return result
}

// https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays/
// 所有奇数长度子数组的和
func sumOddLengthSubarrays(arr []int) int {
	var r int
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j += 2 {
			for k := i; k <= j; k++ {
				r += arr[k]
			}
		}
	}
	return r
}

// https://leetcode-cn.com/problems/create-target-array-in-the-given-order/
// 按既定顺序创建目标数组
func createTargetArray(nums []int, index []int) []int {
	var res = make([]int, len(nums))
	for k, i := range index {
		copy(res[i+1:], res[i:])
		res[i] = nums[k]
	}
	return res
}
