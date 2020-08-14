package main

//字符串的排列
//时间复杂度 O(n!)
//https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/submissions/
//经典回溯法[https://github.com/greyireland/algorithm-pattern/blob/master/advanced_algorithm/backtrack.md]
func permutation(s string) []string {
	ret := make([]string, 0)
	ss := ""
	mark := make([]int, len(s))
	check(ss, sortS(s), &ret, mark)
	return ret
}
func check(ss string, s string, ret *[]string, mark []int) {
	if len(ss) == len(s) {
		*ret = append(*ret, ss)
		return
	}
	for i := 0; i < len(s); i++ {
		if mark[i] < 0 {
			continue
		}
		if i > 0 && mark[i-1] == 0 && s[i] == s[i-1] {
			continue
		}

		mark[i] = -1
		check(ss+string(s[i]), s, ret, mark)
		mark[i] = 0
	}

}
func sortS(s string) string {
	b := []byte(s)
	blen := len(b)

	for i := 0; i < blen-1; i++ {
		for j := 0; j < blen-1-i; j++ {
			if b[j] > b[j+1] {
				temp := b[j]
				b[j] = b[j+1]
				b[j+1] = temp
			}
		}
	}

	return string(b)
}
