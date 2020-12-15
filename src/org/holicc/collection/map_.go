package main

import "unicode"

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

// https://leetcode-cn.com/problems/shortest-completing-word/
// 最短补全词
func shortestCompletingWord(licensePlate string, words []string) string {
	lengthWords := len(words)
	min := 16
	minWord := ""

	//依次遍历取出words里的每个单词
	for i := 0; i < lengthWords; i++ {
		wordLength := len(words[i])
		list := make([]int, 26)

		//如果该单词比当前最短完整词的长度长则舍去
		if wordLength < min {

			//依次将每个单词转换到长度为26的哈希表里(切片做哈希表)
			for j := 0; j < wordLength; j++ {
				list[words[i][j]-'a']++
			}

			//依次遍历比较牌照中的字母是否都在该单词里
			if compare(list, licensePlate) {

				//如果存在则更新最短完整词和最短长度
				min = wordLength
				minWord = words[i]
			}
		}
	}
	return minWord
}

func compare(list []int, licensePlate string) bool {
	for i := 0; i < len(licensePlate); i++ {
		if letter := licensePlate[i]; unicode.IsLetter(rune(letter)) {
			if list[(letter|32)-'a'] <= 0 {
				return false
			}
			list[(letter|32)-'a']--
		}
	}
	return true
}
