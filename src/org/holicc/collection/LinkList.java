package org.holicc.collection;

import org.holicc.model.ListNode;

import java.util.HashSet;
import java.util.LinkedList;
import java.util.Objects;
import java.util.Set;

public class LinkList {


    /**
     * 删除链表的节点
     * 时间复杂度 O(n)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
     */
    public ListNode deleteNode(ListNode head, int val) {
        //
        if (Objects.isNull(head)) return null;
        //
        if (head.val == val) return head.next;
        //
        ListNode tmp = head;
        while (tmp != null) {
            //尾节点
            if (tmp.next != null && tmp.next.val == val) {
                tmp.next = tmp.next.next;
                return head;
            }
            //
            tmp = tmp.next;
        }
        return head;
    }

    /**
     * 两个链表的第一个公共节点
     * 时间复杂度O(n)
     * 空间复杂度O(n)
     * <p>
     * https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
     */
    public ListNode getIntersectionNode_01(ListNode headA, ListNode headB) {
        Set<ListNode> container = new HashSet<>();
        ListNode t = headA;
        while (t != null) {
            container.add(t);
            t = t.next;
        }
        //
        t = headB;
        while (t != null) {
            if (container.contains(t)) {
                return t;
            }
            t = t.next;
        }
        return null;
    }

    /**
     * 两个链表的第一个公共节点
     * 时间复杂度O(n)
     * 空间复杂度O(1)
     * <p>
     * https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
     */
    public ListNode getIntersectionNode_02(ListNode headA, ListNode headB) {
        if (headA == null || headB == null) return null;
        ListNode p = headA;
        ListNode q = headB;
        while (p != q) {
            p = p == null ? headB : p.next;
            q = q == null ? headA : q.next;
        }
        return q;
    }

    /**
     * 合并两个排序的链表
     * 时间复杂度 O(n)
     * 空间复杂度 O(n+m)
     * <p>
     * https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
     */
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode root = new ListNode(0);
        ListNode t = root;
        while (l1 != null && l2 != null) {
            if (l1.val < l2.val) {
                t.next = l1;
                l1 = l1.next;
            } else {
                t.next = l2;
                l2 = l2.next;
            }
            //
            t = t.next;
        }
        t.next = l1 != null ? l1 : l2;
        return root.next;
    }

    /**
     * 反转链表
     * 时间复杂度 O(n)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
     */
    public ListNode reverseList(ListNode head) {
        if (head == null) return null;
        ListNode pre = null;
        ListNode t = head;
        while (t != null) {
            ListNode next = t.next;
            t.next = pre;
            pre = t;
            t = next;
        }
        return pre;
    }

    /**
     * 从尾到头打印链表
     * 时间复杂度 O(n)
     * 空间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
     */
    public int[] reversePrint(ListNode head) {
        if (head == null) return new int[0];
        LinkedList<Integer> queue = new LinkedList<>();
        while (head != null) {
            queue.addFirst(head.val);
            head = head.next;
        }
        //
        int[] r = new int[queue.size()];
        for (int i = 0; i < queue.size(); i++) {
            r[i] = queue.get(i);
        }
        return r;
    }

    /**
     * 链表中倒数第k个节点
     * <p>
     * https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
     */
    public ListNode getKthFromEnd(ListNode head, int k) {
        if (head == null) return null;
        //
        ListNode t = head;
        ListNode p = t;
        //k
        while (k-- > 0) {
            t = t.next;
        }
        //
        while (t != null) {
            t = t.next;
            p = p.next;
        }
        return p;
    }

    /**
     * 左旋转字符串
     * 时间复杂度 O(n)
     * 空间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
     */
    public String reverseLeftWords_01(String s, int n) {
        if (s == null || s.equals("")) return null;
        StringBuilder builder = new StringBuilder(s);
        for (int i = 0; i < n; i++) {
            builder.append(builder.charAt(i));
        }
        return builder.delete(0, n).toString();
    }
}
