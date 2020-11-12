package main

import (
	"math"
	"sort"
	"strconv"
)

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

// https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/
// 矩阵中的幸运数
func luckyNumbers(matrix [][]int) []int {
	var r []int
	for i := 0; i < len(matrix); i++ {
		var c int
		min := math.MaxInt32
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
				c = j
			}
		}
		max := math.MinInt32
		for k := 0; k < len(matrix); k++ {
			if max < matrix[k][c] {
				max = matrix[k][c]
			}
		}
		if max == min {
			r = append(r, min)
		}
	}
	return r
}

// https://leetcode-cn.com/problems/final-prices-with-a-special-discount-in-a-shop/
// 商品折扣后的最终价格
func finalPrices(prices []int) []int {
	stack := make([]int, 0, len(prices))
	for i := 0; i < len(prices); i++ {
		for len(stack) > 0 && prices[i] <= prices[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prices[pop] -= prices[i]
		}
		stack = append(stack, i)
	}
	return prices
}

// https://leetcode-cn.com/problems/find-the-distance-value-between-two-arrays/
// 两个数组间的距离值
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var r int
	for _, v := range arr1 {
		match := true
		for _, c := range arr2 {
			if Abs(v-c) > d {
				match = false
				break
			}
		}
		if match {
			r++
		}
	}
	return r
}

// https://leetcode-cn.com/problems/find-positive-integer-solution-for-a-given-equation/
// 找出给定方程的正整数解
func findSolution(customFunction func(int, int) int, z int) [][]int {
	var res [][]int
	for i := 1; i <= 1000; i++ {
		low, high := 1, 1000
		t1, t2 := customFunction(i, low), customFunction(i, high)
		if z < t1 || z > t2 {
			continue
		}
		for low <= high {
			mid := low + (high-low)/2
			t3 := customFunction(i, mid)
			if z == t3 {
				res = append(res, []int{i, mid})
				break
			} else if z > t3 {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return res
}

// https://leetcode-cn.com/problems/unique-number-of-occurrences/
// 独一无二的出现次数
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	n := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		if _, ok := n[v]; ok {
			return false
		} else {
			n[v] = k
		}
	}
	return true
}

// https://leetcode-cn.com/problems/minimum-subsequence-in-non-increasing-order/
// 非递增顺序的最小子序列
func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	//
	sum := 0
	for _, n := range nums {
		sum += n
	}
	//
	r := make([]int, 0)
	t := 0
	for i := len(nums) - 1; i >= 0; i-- {
		t += nums[i]
		r = append(r, nums[i])
		if t > sum-t {
			return r
		}
	}
	return r
}

// https://leetcode-cn.com/problems/island-perimeter/
// 岛屿的周长
func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				count += 4
				if i > 0 && grid[i-1][j] == 1 {
					count -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					count -= 2
				}
			}
		}
	}
	return count
}

// https://leetcode-cn.com/problems/minimum-value-to-get-positive-step-by-step-sum/
// 逐步求和得到正数的最小值
func minStartValue(nums []int) int {
	sum := 0
	min := 0
	for _, num := range nums {
		sum += num
		if sum < min {
			min = sum
		}
	}
	return 1 - min
}

// https://leetcode-cn.com/problems/distribute-candies/
// 分糖果
func distributeCandies(candies []int) int {
	color := make(map[int]bool)
	for _, val := range candies {
		color[val] = false
	}
	if len(color) > len(candies)/2 {
		return len(candies) / 2
	}
	return len(color)
}

// https://leetcode-cn.com/problems/crawler-log-folder/
// 文件夹操作日志搜集器
func minOperations(logs []string) int {
	s := make([]string, 0)
	for _, o := range logs {
		if o == "../" {
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		} else if o != "./" {
			s = append(s, o)
		}
	}
	return len(s)
}

// https://leetcode-cn.com/problems/pascals-triangle/
// 杨辉三角
func generate(numRows int) [][]int {
	var res [][]int = [][]int{
		{1},
		{1, 1},
	}
	if numRows < 2 {
		return res[0:numRows]
	}

	for i := 2; i < numRows; i++ {
		var row = []int{1}
		for j := 1; j < i; j++ {
			row = append(row, res[i-1][j-1]+res[i-1][j])
		}
		row = append(row, 1)
		res = append(res, row)
	}

	return res
}

// https://leetcode-cn.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
// 	去掉最低工资和最高工资后的工资平均值
func average(salary []int) float64 {
	var max, min = 0, math.MaxInt32
	sum := 0
	for _, v := range salary {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
		sum += v
	}
	return float64(sum-max-min) / float64(len(salary)-2)
}

// https://leetcode-cn.com/problems/minimum-absolute-difference/
// 最小绝对差
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	l := len(arr)
	var result [][]int
	min := arr[l-1] - arr[0]
	for i := 1; i < l; i++ {
		cha := arr[i] - arr[i-1]
		if cha < min {
			min = cha
			result = [][]int{}
			result = append(result, []int{arr[i-1], arr[i]})
		} else if cha == min {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}

// https://leetcode-cn.com/problems/relative-sort-array/
// 数组的相对排序
func relativeSortArray(arr1 []int, arr2 []int) []int {
	hash := make(map[int]int, 0)
	res := make([]int, 0)

	if len(arr2) == 0 {
		sort.Ints(arr1)
		return arr1
	}

	for _, v := range arr1 {
		hash[v]++
	}
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < hash[arr2[i]]; j++ {
			res = append(res, arr2[i])
		}
		hash[arr2[i]] = 0
	}

	temp := make([]int, 0)
	for k, v := range hash {
		for v > 0 {
			temp = append(temp, k)
			v--
		}
	}
	sort.Ints(temp)

	res = append(res, temp...)
	return res
}

// https://leetcode-cn.com/problems/reshape-the-matrix/solution/golang-yi-ceng-for15xing-dai-ma-si-lu-qing-xi-ban-/
// 重塑矩阵 取余取模，定位索引
func matrixReshape(nums [][]int, r int, c int) [][]int {
	row, col := len(nums), len(nums[0])
	if row*col != r*c {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, c)
	}

	for i := 0; i < r*c; i++ {
		res[i/c][i%c] = nums[i/col][i%col]
	}
	return res
}

// https://leetcode-cn.com/problems/toeplitz-matrix/
// 托普利茨矩阵
func isToeplitzMatrix(matrix [][]int) bool {
	hash := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if _, ok := hash[i-j]; ok {
				if hash[i-j] != matrix[i][j] {
					return false
				}
			} else {
				hash[i-j] = matrix[i][j]
			}
		}
	}
	return true
}

// https://leetcode-cn.com/problems/projection-area-of-3d-shapes/
// 三维形体投影面积
func projectionArea(grid [][]int) int {
	N := len(grid)
	result := 0
	for row := 0; row < N; row++ {
		rowMax, colMax := 0, 0
		for col := 0; col < N; col++ {
			if grid[row][col] != 0 {
				result++
			}
			if grid[row][col] > rowMax {
				rowMax = grid[row][col]
			}
			if grid[col][row] > colMax {
				colMax = grid[col][row]
			}
		}
		result += rowMax + colMax
	}
	return result
}

// https://leetcode-cn.com/problems/next-greater-element-i/solution/golangdan-diao-zhan-fa-by-bloodborne/
// 下一个更大元素 I
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var stack []int
	for _, v := range nums2 {
		for len(stack) != 0 && v > stack[len(stack)-1] {
			// 发现有更大的数字，给其下一个更大数字赋值
			m[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	for k, v := range nums1 {
		if value, ok := m[v]; ok {
			nums1[k] = value
		} else {
			nums1[k] = -1
		}
	}
	return nums1
}

// https://leetcode-cn.com/problems/build-an-array-with-stack-operations/
// 用栈操作构建数组
func buildArray(target []int, n int) []string {
	res := make([]string, 0)
	num := 0
	for i := 1; i <= n; i++ {
		if i > target[len(target)-1] {
			return res
		} else if i == target[num] {
			res = append(res, "Push")
			num++
		} else {
			res = append(res, "Push")
			res = append(res, "Pop")
		}
	}
	return res
}

// https://leetcode-cn.com/problems/majority-element/
// 多数元素
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v >= len(nums)/2 {
			return k
		}
	}
	return -1
}

// https://leetcode-cn.com/problems/matrix-cells-in-distance-order/
// 距离顺序排列矩阵单元格
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	result := make([][]int, R*C)
	k := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			result[k] = []int{r, c}
			k++
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return Abs(result[i][0]-r0)+Abs(result[i][1]-c0) < Abs(result[j][0]-r0)+Abs(result[j][1]-c0)
	})
	return result
}

func main() {
	kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, 3)
}

// https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/
// 方阵中战斗力最弱的 K 行
func kWeakestRows(mat [][]int, k int) []int {
	ans := make([]int, 0, k)
	m := make(map[int][]int)
	for index, line := range mat {
		num := 0
		for _, v := range line {
			if 1 == v {
				num++
			}
		}
		m[num] = append(m[num], index)
	}
	for i := 0; i <= len(mat[0]); i++ {
		ans = append(ans, m[i]...)
		if len(ans) >= k {
			ans = ans[:k]
			break
		}
	}

	return ans
}

// https://leetcode-cn.com/problems/distribute-candies-to-people/
// 分糖果 II
func distributeCandies2(candies int, num_people int) []int {
	r := make([]int, num_people)
	for i, l := 0, 0; candies > 0; i++ {
		if i >= len(r) {
			i = 0
			l++
		}
		c := i + (num_people * l) + 1
		if c > candies {
			r[i] += candies
		} else {
			r[i] += c
		}
		candies -= c
	}
	return r
}

// https://leetcode-cn.com/problems/largest-triangle-area/
// 最大三角形面积
func largestTriangleArea(points [][]int) float64 {
	res := 0
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				temp := points[i][0]*points[j][1] + points[j][0]*points[k][1] + points[k][0]*points[i][1] - points[j][1]*points[k][0] - points[i][1]*points[j][0] - points[i][0]*points[k][1]
				if abs(temp) > res {
					res = abs(temp)
				}
			}
		}
	}
	return float64(res) / 2
}

func abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

// https://leetcode-cn.com/problems/element-appearing-more-than-25-in-sorted-array/
// 有序数组中出现次数超过25%的元素
func findSpecialInteger(arr []int) int {

	var n = (len(arr) + 4) / 4

	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[n-1+i] {
			return arr[i]
		}
	}
	return -1
}
