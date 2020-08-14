package main

import "strconv"

func main() {
	println(movingCount(3, 1, 0))
}

//机器人的运动范围
//https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
//时间复杂度 O(n*m)
//空间复杂度 O(n*m)
//DFS遍历就行了
func movingCount(m int, n int, k int) int {
	visit := make2DAry(m, n)
	moving(visit, 0, 0, k)
	r := 0
	for _, v := range visit {
		for _, c := range v {
			r += c
		}
	}
	return r
}

func moving(visit [][]int, i, j, k int) {
	//
	if i >= len(visit) || j >= len(visit[0]) {
		return
	}
	//
	//
	if visit[i][j] == 0 && (calculate(i)+calculate(j)) <= k {
		visit[i][j] = 1
		//
		moving(visit, i+1, j, k)
		//
		moving(visit, i, j+1, k)
	}
}

func calculate(i int) (sum int) {
	for _, c := range strconv.Itoa(i) {
		sum += int(c - '0')
	}
	return
}

func make2DAry(m, n int) [][]int {
	ary := make([][]int, m)
	for i := 0; i < m; i++ {
		ary[i] = make([]int, n)
	}
	return ary
}
