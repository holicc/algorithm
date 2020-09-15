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
