package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(restoreString("aiohn", []int{3, 1, 4, 2, 0}))
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

// https://leetcode-cn.com/problems/decompress-run-length-encoded-list/
// 解压缩编码列表
func decompressRLElist(nums []int) []int {
	r := make([]int, 0)
	for i := 1; i < len(nums); i += 2 {
		t := make([]int, 0)
		for j := 0; j < nums[i-1]; j++ {
			t = append(t, nums[i])
		}
		r = append(r, t...)
	}
	return r
}

// https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/
// 将数字变成 0 的操作次数
func numberOfSteps(num int) int {
	var r int
	for num > 0 {
		if (num & 1) == 1 {
			num -= 1
		} else {
			num /= 2
		}
		r++
	}
	return r
}

// https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/
// 有多少小于当前数字的数字
func smallerNumbersThanCurrent(nums []int) []int {
	r := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		t := nums[i]
		var c int
		for j := 0; j < len(nums); j++ {
			if nums[j] < t {
				c++
			}
		}
		r[i] = c
	}
	return r
}

// https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits/
// 统计位数为偶数的数字
func findNumbers(nums []int) int {
	var r int
	for i := range nums {
		itoa := strconv.Itoa(nums[i])
		if (len(itoa) & 1) == 0 {
			r++
		}
	}
	return r
}

// https://leetcode-cn.com/problems/minimum-time-visiting-all-points/
// 访问所有点的最小时间
func minTimeToVisitAllPoints(points [][]int) int {
	var r int
	for i := 1; i < len(points); i++ {
		pre := points[i-1]
		cur := points[i]
		x := int(math.Abs(float64(cur[0] - pre[0])))
		y := int(math.Abs(float64(cur[1] - pre[1])))
		if x > y {
			r += x
		} else {
			r += y
		}
	}
	return r
}

// https://leetcode-cn.com/problems/number-of-students-doing-homework-at-a-given-time/
// 在既定时间做作业的学生人数
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var r int
	for i := range startTime {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			r++
		}
	}
	return r
}

// https://leetcode-cn.com/problems/shuffle-string/
// 重新排列字符串
func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i := range s {
		res[indices[i]] = s[i]
	}
	return string(res)
}

// https://leetcode-cn.com/problems/count-good-triplets/
// 统计好三元组
func countGoodTriplets(arr []int, a int, b int, c int) int {
	var r int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if Abs(arr[i]-arr[j]) <= a &&
					Abs(arr[j]-arr[k]) <= b &&
					Abs(arr[i]-arr[k]) <= c {
					r++
				}
			}
		}
	}
	return r
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

// https://leetcode-cn.com/problems/maximum-product-of-two-elements-in-an-array/
// 数组中两元素的最大乘积
func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}

// https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/
// 将每个元素替换为右侧最大元素
func replaceElements(arr []int) []int {
	m := -1
	for i := len(arr) - 1; i > 0; i-- {
		if m < arr[i] {
			arr[i], m = m, arr[i]
		} else {
			arr[i] = m
		}
	}
	return arr
}

// https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix/
// 统计有序矩阵中的负数
func countNegatives(grid [][]int) int {
	var r int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] < 0 {
				r++
			}
		}
	}
	return r
}

// https://leetcode-cn.com/problems/cells-with-odd-values-in-a-matrix/
// 奇数值单元格的数目
func oddCells(n int, m int, indices [][]int) int {

	var r int
	rows := make([]int, n)
	cols := make([]int, m)
	for i := range indices {
		arr := indices[i]
		row, col := arr[0], arr[1]
		rows[row]++
		cols[col]++
	}
	var rowNum int
	for _, v := range rows {
		if (v & 1) == 1 {
			r += m
			rowNum++
		}
	}
	for _, v := range cols {
		if (v & 1) == 1 {
			r += n - 2*rowNum
		}
	}
	return r
}

// https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays/
// 通过翻转子数组使两个数组相等
func canBeEqual(target []int, arr []int) bool {
	sort.Ints(arr)
	sort.Ints(target)
	for i := range target {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}

// https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero/
// 和为零的N个唯一整数
func sumZero(n int) []int {
	var s int
	t := n
	r := make([]int, 0, n)
	for t > 1 {
		r = append(r, -t)
		s += -t
		t--
	}
	r = append(r, -s)
	return r
}

// https://leetcode-cn.com/problems/can-make-arithmetic-progression-from-sequence/
// 判断能否形成等差数列
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	d := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != d {
			return false
		}
	}
	return true
}
