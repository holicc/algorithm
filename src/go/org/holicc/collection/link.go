package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/reverse-linked-list/
// 反转链表
func reverseList(head *ListNode) *ListNode {
	var t *ListNode
	var p *ListNode
	for head != nil {
		t = head.Next
		head.Next = p
		p = head
		head = t
	}
	return p
}

// https://leetcode-cn.com/problems/middle-of-the-linked-list/
// 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// https://leetcode-cn.com/problems/three-consecutive-odds/
// 存在连续三个奇数的数组
func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i < len(arr)-2; i++ {
		if ((arr[i] & 1) & (arr[i+1] & 1) & (arr[i+2] & 1)) == 1 {
			return true
		}
	}
	return false
}

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var r = &ListNode{0, head}
	t := r
	for t.Next != nil && t.Next.Next != nil {
		n1 := t.Next
		n2 := t.Next.Next
		t.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		t = n1
	}
	return r.Next
}

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{
		Val:  -1,
		Next: nil,
	}
	rr := r
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			r.Next = l2
			l2 = l2.Next
		} else if l1.Val <= l2.Val {
			r.Next = l1
			l1 = l1.Next
		}
		r = r.Next
	}
	if l1 != nil {
		r.Next = l1
	}
	if l2 != nil {
		r.Next = l2
	}
	return rr.Next
}

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
// 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	tmpa, tmpb := headA, headB
	for tmpa != tmpb {
		if tmpa == nil {
			tmpa = headB
		} else {
			tmpa = tmpa.Next
		}

		if tmpb == nil {
			tmpb = headA
		} else {
			tmpb = tmpb.Next
		}
	}

	return tmpa
}
