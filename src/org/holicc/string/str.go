package main

import (
	"sort"
	"strconv"
	"strings"
)

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

//表示数值的字符串
//https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
type State int
type CharType int

const (
	STATE_INITIAL State = iota
	STATE_INT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_WITHOUT_INT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_NUMBER
	STATE_END
)

const (
	CHAR_NUMBER CharType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_SPACE
	CHAR_ILLEGAL
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	case ' ':
		return CHAR_SPACE
	default:
		return CHAR_ILLEGAL
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		STATE_INITIAL: map[CharType]State{
			CHAR_SPACE:  STATE_INITIAL,
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
			CHAR_SIGN:   STATE_INT_SIGN,
		},
		STATE_INT_SIGN: map[CharType]State{
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
		},
		STATE_INTEGER: map[CharType]State{
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_EXP:    STATE_EXP,
			CHAR_POINT:  STATE_POINT,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT_WITHOUT_INT: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
		},
		STATE_FRACTION: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_EXP: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SIGN:   STATE_EXP_SIGN,
		},
		STATE_EXP_SIGN: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
		},
		STATE_EXP_NUMBER: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SPACE:  STATE_END,
		},
		STATE_END: map[CharType]State{
			CHAR_SPACE: STATE_END,
		},
	}
	state := STATE_INITIAL
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == STATE_INTEGER || state == STATE_POINT || state == STATE_FRACTION || state == STATE_EXP_NUMBER || state == STATE_END
}

//https://leetcode-cn.com/problems/nGK0Fy/
//速算机器人
func calculate(s string) int {
	var x, y = 1, 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "A" {
			x = 2*x + y
		}
		if string(s[i]) == "B" {
			y = 2*y + x
		}
	}
	return x + y
}

// https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
// 整数的各位积和之差
func subtractProductAndSum(n int) int {
	m := 1
	s := 0
	for n > 0 {
		i := n % 10
		m *= i
		s += i
		n /= 10
	}
	return m - s
}

// https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/
// 分割平衡字符串
func balancedStringSplit(s string) int {
	var (
		r int
		c int
	)
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			c++
		} else {
			c--
		}
		if c == 0 {
			r++
		}
	}
	return r
}

// https://leetcode-cn.com/problems/destination-city/
// 旅行终点站
func destCity(paths [][]string) string {
	m := make(map[string]string)
	for i := range paths {
		path := paths[i]
		m[path[0]] = path[1]
	}
	//
	for i := range paths {
		path := paths[i]
		if _, ok := m[path[0]]; !ok {
			return path[0]
		}
		if _, ok := m[path[1]]; !ok {
			return path[1]
		}
	}
	return ""
}

// https://leetcode-cn.com/problems/maximum-69-number/
// 组成的最大数字
func maximum69Number(num int) int {
	bytes := []byte(strconv.Itoa(num))
	for i := range bytes {
		if bytes[i] == '6' {
			bytes[i] = '9'
			break
		}
	}
	r, _ := strconv.Atoi(string(bytes))
	return r
}

// https://leetcode-cn.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
// 解码字母到整数映射
func freqAlphabets(s string) string {
	n, ans := len(s), []rune{}
	for i := 0; i < n; i++ {
		if i+2 < n && s[i+2] == '#' {
			ans = append(ans, get(s[i:i+2]))
			i += 2
		} else {
			ans = append(ans, get(s[i:i+1]))
		}
	}
	return string(ans)
}

func get(s string) rune {
	r, _ := strconv.Atoi(s)
	return rune(r + 96)
}

// https://leetcode-cn.com/problems/increasing-decreasing-string/
// 上升下降字符串
func sortString(s string) string {
	var table [26]byte
	var r strings.Builder
	for _, v := range s {
		table[v-'a'] += 1
	}
	//
	len_ := len(s)
	for len_ > 0 {
		for i := 0; i < 26; i++ {
			if table[i] > 0 {
				r.WriteByte('a' + byte(i))
				table[i]--
				len_--
			}
		}
		for i := 25; i >= 0; i-- {
			if table[i] > 0 {
				r.WriteByte('a' + byte(i))
				table[i]--
				len_--
			}
		}
	}
	return r.String()
}

// https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/
// 最长特殊序列 Ⅰ
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	} else {
		return len(b)
	}
}

// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/
// 拼写单词
func countCharacters(words []string, chars string) int {
	var r int
	//
	m := make([]int, 26)
	for _, c := range chars {
		m[c-'a']++
	}
	//
	for _, w := range words {
		match := true
		t := make([]int, 26)
		for _, c := range w {
			t[c-'a']++
		}
		//
		for _, c := range w {
			if t[c-'a'] > m[c-'a'] {
				match = false
				break
			}
		}
		if match {
			r += len(w)
		}
	}
	return r
}

// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/
// 删除字符串中的所有相邻重复项
func removeDuplicates(S string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}

// https://leetcode-cn.com/problems/baseball-game/
// 棒球比赛
func calPoints(ops []string) int {
	var r int
	stack := make([]int, 0)
	for _, c := range ops {
		switch c {
		case "+":
			if len(stack) >= 1 {
				stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
			} else if len(stack) > 0 {
				stack = append(stack, stack[len(stack)-1])
			}
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "C":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			t, _ := strconv.Atoi(c)
			stack = append(stack, t)
		}
	}
	//
	for _, c := range stack {
		r += c
	}
	return r
}

// https://leetcode-cn.com/problems/remove-palindromic-subsequences/
// 删除回文子序列
func removePalindromeSub(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	if s == reverseString(s) {
		return 1
	}
	return 2
}

func reverseString(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// https://leetcode-cn.com/problems/number-of-lines-to-write-string/
// 写字符串需要的行数
func numberOfLines(widths []int, S string) []int {
	res := make([]int, 2)
	var length, row int = 0, 1 //length为当前行的当前长度，row为当前行数
	for i := 0; i < len(S); i++ {
		if length+widths[S[i]-'a'] > 100 { //如果当前长度加上下一个字符长度超过100，说明当前行剩下的长度
			row++ //不足以放下一个字符，所以重开一行，放下一个字符，length变为新一行当前的长度
			length = widths[S[i]-'a']
		} else {
			length += widths[S[i]-'a']
		}
	}
	res[0] = row    //行数
	res[1] = length //最后一行长度
	return res
}

// https://leetcode-cn.com/problems/largest-substring-between-two-equal-characters/
// 两个相同字符之间的最长子字符串
func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]] = i
	}
	//
	var r = -1
	for i := range s {
		if r < m[s[i]]-i {
			r = m[s[i]] - i
		}
	}
	return r - 1
}

// https://leetcode-cn.com/problems/find-the-difference/
// 找不同
func findTheDifference(s string, t string) byte {
	hash := make(map[byte]int)
	var res byte

	for k, _ := range t {
		hash[t[k]]++
	}

	for k, _ := range s {
		hash[s[k]]--
	}

	for k, v := range hash {
		if v == 1 {
			res = k
		}
	}
	return res
}

// https://leetcode-cn.com/problems/count-binary-substrings/
// 计数二进制子串
func countBinarySubstrings(s string) int {
	counts := []int{}
	ptr, n := 0, len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		counts = append(counts, count)
	}
	ans := 0
	for i := 1; i < len(counts); i++ {
		ans += min(counts[i], counts[i-1])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// https://leetcode-cn.com/problems/occurrences-after-bigram/
// Bigram 分词
func findOcurrences(text string, first string, second string) []string {
	var res []string
	s := strings.Split(text, " ")
	for i := 0; i < len(s)-2; i++ {
		if s[i] == first && s[i+1] == second {
			res = append(res, s[i+2])
		}
	}
	return res
}

// https://leetcode-cn.com/problems/compare-strings-by-frequency-of-the-smallest-character/submissions/
// 比较字符串最小字母出现频次
func numSmallerByFrequency(queries []string, words []string) (ans []int) {
	for _, qv := range queries {
		c := 0
		for _, wv := range words {
			if do(qv, wv) == true {
				c++
			}
		}
		ans = append(ans, c)
	}
	return ans
}
func do(p, q string) (ans bool) {
	return count(p) < count(q)
}
func count(s string) (ans int) {
	ar := []byte(s)
	k := ar[0]
	for _, v := range ar {
		if v < k {
			k = v
		}
	}
	return strings.Count(s, string(k))
}

// https://leetcode-cn.com/problems/string-matching-in-an-array/submissions/
// 数组中的字符串匹配
func stringMatching(words []string) []string {
	d := make(map[string]bool, len(words))
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	for i, n := 0, len(words); i < n; i++ {
		for j := i + 1; j < n; j++ {
			if strings.Contains(words[j], words[i]) {
				d[words[i]] = false
			}
		}
	}
	ans, i := make([]string, len(d)), 0
	for w := range d {
		ans[i] = w
		i++
	}
	return ans
}

// https://leetcode-cn.com/problems/goat-latin/
// 山羊拉丁文
func toGoatLatin(S string) string {
	var yuanYinMap = map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	sliceWorlds := strings.Split(S, " ")
	result := make([]string, len(sliceWorlds))
	i := 0
	for _, world := range sliceWorlds {
		tmp := []byte(world)
		if _, ok := yuanYinMap[tmp[0]]; ok {
			result[i] = world + "ma"
		} else {
			result[i] = world[1:] + world[0:1] + "ma"
		}
		result[i] = result[i] + strings.Repeat("a", i+1)
		i++
	}

	return strings.Join(result, " ")
}

// https://leetcode-cn.com/problems/palindrome-number/
// 回文数
func isPalindrome(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

// https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/
// 字符串的最大公因子
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	gcd := func(a, b int) int {
		r := a % b
		for r != 0 {
			a = b
			b = r
			r = a % b
		}
		return b
	}
	return str1[0:gcd(len(str1), len(str2))]
}

// https://leetcode-cn.com/problems/consecutive-characters/
func maxPower(s string) int {
	var (
		r int = 1
		x int = 1
	)
	for i := 1; i < len(s)-1; i++ {
		if s[i] == s[i-1] {
			r++
			if r > x {
				x = r
			}
		} else {
			r = 1
		}
	}
	return x
}
