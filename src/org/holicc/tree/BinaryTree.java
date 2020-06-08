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
}
