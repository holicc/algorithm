package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	isSubStructure(&TreeNode{
		Val: 1,
	}, &TreeNode{
		Val: 1,
	})
}

//二叉搜索树的后序遍历序列
//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/submissions/
//时间复杂度 O(N^2)O(N)每次调用 recur(i,j)recur(i,j) 减去一个根节点，因此递归占用 O(N)O(N) ；最差情况下（即当树退化为链表），每轮递归都需遍历树所有节点，占用 O(N)O(N) 。
//空间复杂度 O(N)O(N) ： 最差情况下（即当树退化为链表），递归深度将达到 NN 。
func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	idx := i
	for postorder[idx] < postorder[j] {
		idx++
	}
	mid := idx
	for postorder[idx] > postorder[j] {
		idx++
	}
	return idx == j && recur(postorder, i, mid-1) && recur(postorder, mid, j-1)
}

//二叉搜索树的后序遍历序列
//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/solution/mian-shi-ti-33-er-cha-sou-suo-shu-de-hou-xu-bian-6/
//时间复杂度 O(N) 遍历 postorderpostorder 所有节点，各节点均入栈 / 出栈一次，使用 O(N)O(N) 时间。
//空间复杂度 O(N) 最差情况下，单调栈 stackstack 存储所有节点，使用 O(N)O(N) 额外空间。
func st(postorder []int) bool {
	l := -1
	stack := make([]int, 0)
	//
	root := math.MaxInt32
	//
	for i := len(postorder) - 1; i >= 0; i-- {
		v := postorder[i]
		if v > root {
			return false
		}
		//
		for l >= 0 && stack[l] > v {
			root = stack[l]
			stack = stack[:l]
			l--
		}
		//
		stack = append(stack, v)
		l++
	}
	return true
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (r(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func r(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil || a.Val != b.Val {
		return false
	}
	return r(a.Left, b.Left) && r(a.Right, b.Right)
}
