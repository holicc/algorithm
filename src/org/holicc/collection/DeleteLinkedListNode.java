package org.holicc.collection;

import org.holicc.model.ListNode;

import java.util.Objects;

public class DeleteLinkedListNode {


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

}
