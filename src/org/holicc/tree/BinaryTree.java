package org.holicc.tree;

import org.holicc.model.TreeNode;

import java.util.*;


public class BinaryTree {

    /**
     * 二叉搜索树的最近公共祖先
     * <p>
     * 时间复杂度 O(n log(n)) 树高
     * 空间复杂度 O(n log(n)) 递归使用的栈空间
     * <p>
     * https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/
     * https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/submissions/
     */
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        //
        if (root == null) return null;
        if (root == p) return p;
        if (root == q) return q;
        //
        TreeNode left = lowestCommonAncestor(root.left, p, q);
        TreeNode right = lowestCommonAncestor(root.right, p, q);
        if (left != null && right != null) return root;
        return left == null ? right : left;
    }

    /**
     * 从上到下打印二叉树 II
     * <p>
     * https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
     */
    public List<List<Integer>> levelOrder(TreeNode root) {
        Queue<TreeNode> queue = new LinkedList<>();
        List<List<Integer>> res = new ArrayList<>();
        if (root != null) queue.add(root);
        while (!queue.isEmpty()) {
            List<Integer> tmp = new ArrayList<>();
            for (int i = queue.size(); i > 0; i--) {
                TreeNode node = queue.poll();
                tmp.add(node.val);
                if (node.left != null) queue.add(node.left);
                if (node.right != null) queue.add(node.right);
            }
            res.add(tmp);
        }
        return res;
    }

    /**
     * 二叉搜索树的第k大节点
     * 时间复杂度 O(n log(n))
     * <p>
     * https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
     */
    public int kthLargest(TreeNode root, int k) {
        Stack<TreeNode> stack = new Stack<>();
        while (root != null || !stack.isEmpty()) {
            if (root != null) {
                stack.push(root);
                root = root.right;
            } else {
                TreeNode pop = stack.pop();
                if (--k == 0) return pop.val;
                if (pop.left != null) {
                    root = pop.left;
                }
            }
        }
        return -1;
    }

    /**
     * 二叉树的深度
     * 时间复杂度 O(log(n))
     * <p>
     * https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
     */
    public int maxDepth(TreeNode root) {
        if (root == null) return 0;
        return Math.max(deep(root.left, 1), deep(root.right, 1));
    }

    private int deep(TreeNode root, int i) {
        if (root == null) return i;
        return Math.max(deep(root.left, i + 1), deep(root.right, i + 1));
    }

    /**
     * 二叉树的镜像
     * 时间复杂度 O(log(n))
     * <p>
     * https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
     */
    public TreeNode mirrorTree_01(TreeNode root) {
        if (root == null) return null;
        TreeNode tmp = root.left;
        root.left = mirrorTree_01(root.right);
        root.right = mirrorTree_01(tmp);
        return root;
    }

    /**
     * 二叉树的镜像
     * 时间复杂度 O(log(n))
     * <p>
     * https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
     */
    public TreeNode mirrorTree_02(TreeNode root) {
        Stack<TreeNode> stack = new Stack<>();
        if (root != null) stack.push(root);
        while (!stack.isEmpty()) {
            TreeNode pop = stack.pop();
            if (pop.left != null) stack.push(pop.left);
            if (pop.right != null) stack.push(pop.right);
            TreeNode t = pop.left;
            pop.left = pop.right;
            pop.right = t;
        }
        return root;
    }

    /**
     * 平衡二叉树
     * 时间复杂度 O(n log(n))
     * 空间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
     */
    public boolean isBalanced(TreeNode root) {
        if (root == null) return true;
        return Math.abs(deep(root.left, 1) - deep(root.right, 1)) <= 1
                && isBalanced(root.left)
                && isBalanced(root.right);
    }

    /**
     * 对称的二叉树
     * 时间复杂度 O(n log(n))
     * 空间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
     */
    public boolean isSymmetric(TreeNode root) {
        if (root == null) return true;
        //
        return isSymmetric(root.left, root.right);
    }

    private boolean isSymmetric(TreeNode left, TreeNode right) {
        if (left == null && right == null) return true;
        if (left == null || right == null) return false;
        if (left.val != right.val) return false;
        return isSymmetric(left.left, right.right)
                && isSymmetric(left.right, right.left);
    }

    /**
     * 重建二叉树
     * <p>
     * https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
     */
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        if (preorder == null || inorder == null) return null;
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < preorder.length; i++) {
            map.put(inorder[i], i);
        }
        return buildTree(preorder, 0, preorder.length - 1,
                0, inorder.length - 1, map);
    }

    private TreeNode buildTree(int[] preorder, int preorderStart, int preorderEnd, int inorderStart, int inorderEnd, Map<Integer, Integer> indexMap) {
        if (preorderStart > preorderEnd) return null;
        int rootVal = preorder[preorderStart];
        TreeNode root = new TreeNode(rootVal);
        if (preorderStart != preorderEnd) {
            Integer rootIndex = indexMap.get(rootVal);
            int leftNodes = rootIndex - inorderStart;
            int rightNodes = inorderEnd - rootIndex;
            TreeNode leftSubtree = buildTree(preorder, preorderStart + 1, preorderStart + leftNodes, inorderStart, rootIndex - 1, indexMap);
            TreeNode rightSubtree = buildTree(preorder, preorderEnd - rightNodes + 1, preorderEnd, rootIndex + 1, inorderEnd, indexMap);
            root.left = leftSubtree;
            root.right = rightSubtree;
        }
        return root;
    }

    /**
     * 从上到下打印二叉树
     * BFS
     * <p>
     * https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
     */
    public int[] levelOrder_01(TreeNode root) {
        if (root == null) return new int[0];
        List<Integer> r = new ArrayList<>();
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            TreeNode poll = queue.poll();
            if (poll.left != null) queue.add(poll.left);
            if (poll.right != null) queue.add(poll.right);
            r.add(poll.val);
        }
        //
        return r.stream().mapToInt(i -> i).toArray();
    }

    /**
     * 二叉搜索树与双向链表
     * <p>
     * https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
     */
    TreeNode pre, head;

    public TreeNode treeToDoublyList(TreeNode root) {
        if (root == null) return null;
        dfs(root);
        head.left = pre;
        pre.right = head;
        return head;
    }

    void dfs(TreeNode cur) {
        if (cur == null) return;
        dfs(cur.left);
        if (pre != null) pre.right = cur;
        else head = cur;
        cur.left = pre;
        pre = cur;
        dfs(cur.right);
    }

    /**
     * 二叉树中和为某一值的路径
     * <p>
     * https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
     */
    LinkedList<List<Integer>> res = new LinkedList<>();
    LinkedList<Integer> path = new LinkedList<>();
    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        recur(root, sum);
        return res;
    }
    void recur(TreeNode root, int tar) {
        if(root == null) return;
        path.add(root.val);
        tar -= root.val;
        if(tar == 0 && root.left == null && root.right == null)
            res.add(new LinkedList(path));
        recur(root.left, tar);
        recur(root.right, tar);
        path.removeLast();
    }
}
