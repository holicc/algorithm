package org.holicc.collection;

import org.holicc.model.ListNode;

import java.util.HashSet;
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

}
