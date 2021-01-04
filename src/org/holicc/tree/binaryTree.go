package main

import (
	"math"
	"strconv"
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

func increasingBST(root *TreeNode) *TreeNode {
	var nums = make([]int, 0)
	dfs(root, &nums)
	r := &TreeNode{
		Val: nums[0],
	}
	t := r
	for _, v := range nums[1:] {
		next := &TreeNode{
			Val: v,
		}
		t.Right = next
		t = t.Right
	}
	return r
}

func dfs(root *TreeNode, num *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, num)
	*num = append(*num, root.Val)
	dfs(root.Right, num)
}

// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
// 二叉树的层平均值
func averageOfLevels(root *TreeNode) []float64 {
	var r []float64
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		var s float64
		t := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			v := queue[i]
			s += float64(v.Val)
			if v.Left != nil {
				t = append(t, v.Left)
			}
			if v.Right != nil {
				t = append(t, v.Right)
			}
		}
		r = append(r, s/float64(len(queue)))
		queue = t
	}
	return r
}

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
// 二叉树的层次遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	var r [][]int
	for len(queue) > 0 {
		rr := make([]int, 0)
		t := make([]*TreeNode, 0)
		for _, n := range queue {
			rr = append(rr, n.Val)
			if n.Left != nil {
				t = append(t, n.Left)
			}
			if n.Right != nil {
				t = append(t, n.Right)
			}
		}
		r = append(r, rr)
		queue = t
	}
	for i := 0; i < len(r)/2; i++ {
		t := r[i]
		r[i] = r[len(r)-1-i]
		r[len(r)-1-i] = t
	}
	return r
}

// https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/solution/
// 从根到叶的二进制数之和
func sumRootToLeaf(root *TreeNode) int {
	return f(root, 0)
}

func f(root *TreeNode, r int) int {
	if root != nil {
		r = r*2 + root.Val
		if root.Left == nil && root.Right == nil {
			return r
		}
		return f(root.Left, r) + f(root.Right, r)
	}
	return 0
}

// https://leetcode-cn.com/problems/trim-a-binary-search-tree/
// 修剪二叉搜索树
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}

// https://leetcode-cn.com/problems/binary-tree-paths/
// 二叉树的所有路径
var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	constructPaths(root, "")
	return paths
}

func constructPaths(root *TreeNode, path string) {
	if root != nil {
		pathSB := path
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, pathSB)
		} else {
			pathSB += "->"
			constructPaths(root.Left, pathSB)
			constructPaths(root.Right, pathSB)
		}
	}
}

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}

// https://leetcode-cn.com/problems/leaf-similar-trees/
// 叶子相似的树
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1, l2 := []int{}, []int{}
	dfs_0(root1, &l1)
	dfs_0(root2, &l2)
	if len(l1) != len(l2) {
		return false
	}
	for i, _ := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func dfs_0(root *TreeNode, leafs *[]int) {
	if root != nil {
		dfs(root.Left, leafs)
		if root.Left == nil && root.Right == nil {
			*leafs = append(*leafs, root.Val)
		}
		dfs(root.Right, leafs)
	}
}

// https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
// 二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

// https://leetcode-cn.com/problems/same-tree/
// 相同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// https://leetcode-cn.com/problems/binary-tree-tilt/
// 二叉树的坡度
func findTilt(root *TreeNode) int {
	var (
		i   int
		dfs func(root *TreeNode) int
	)
	abs := func(i int) int {
		if i < 0 {
			return -i
		} else {
			return i
		}
	}
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		i += abs(l - r)
		return root.Val + l + r
	}
	dfs(root)
	return i
}

// https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/
// 两数之和 IV - 输入 BST
func findTarget(root *TreeNode, k int) (result bool) {
	path := findBSTPath(root, map[int]bool{})

	for key, _ := range path {
		if path[k-key] && k-key != key {
			return true
		}
	}

	return
}

func findBSTPath(root *TreeNode, path map[int]bool) map[int]bool {
	if root != nil {
		path = findBSTPath(root.Left, path)
		path[root.Val] = true
		path = findBSTPath(root.Right, path)
	}

	return path
}
