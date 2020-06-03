package org.holicc.tree;

import org.holicc.model.TreeNode;

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

}
