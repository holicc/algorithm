package main

import "strconv"

func main() {
	println(movingCount(3, 1, 0))
}

//https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
//矩阵中的路径
//时间复杂度 O(3^KMN)O(3KMN) ： 最差情况下，需要遍历矩阵中长度为 KK 字符串的所有方案，时间复杂度为 O(3^K)O(3K )；矩阵中共有 MNMN 个起点，时间复杂度为 O(MN)O(MN) 。
//方案数计算： 设字符串长度为 KK ，搜索中每个字符有上、下、左、右四个方向可以选择，舍弃回头（上个字符）的方向，剩下 33 种选择，因此方案数的复杂度为 O(3^K)O(3K) 。
//空间复杂度 O(K) ： 搜索过程中的递归深度不超过 KK ，因此系统因函数调用累计使用的栈空间占用 O(K)O(K) （因为函数返回后，系统调用的栈空间会释放）。最坏情况下 K = MNK=MN ，递归深度为 MNMN ，此时系统栈使用 O(MN)O(MN) 的额外空间。
func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if bt(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func bt(board [][]byte, x, y, i int, word string) bool {
	if x >= len(board) || y >= len(board[0]) || x < 0 || y < 0 || board[x][y] != word[i] {
		return false
	}
	if len(word)-1 == i {
		return true
	}
	tmp := board[x][y]
	board[x][y] = '/'
	r := bt(board, x+1, y, i+1, word) || bt(board, x-1, y, i+1, word) ||
		bt(board, x, y+1, i+1, word) || bt(board, x, y-1, i+1, word)
	//
	board[x][y] = tmp
	return r
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
