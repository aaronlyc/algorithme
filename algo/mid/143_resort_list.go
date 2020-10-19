package mid

import (
	s "example.com/algorithms/algo/struct"
)

func reorderList(head *s.ListNode) {
	if head == nil {
		return
	}
	midNode := findMidNodeFront(head)
	right := reverseList(midNode.Next)
	midNode.Next = nil
	head = mergeListFromLeft(head, right)
}

func findMidNodeFront(head *s.ListNode) *s.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func reverseList(head *s.ListNode) *s.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *s.ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre, head = head, tmp
	}
	return pre
}

//func mergeListFromLeft(head, l2 *s.ListNode) {
//if l2 == nil {
//	return
//}
//l1 := head
//for l1 != nil {
//	tmp1, tmp2 := l1.Next, l2.Next
//	l1.Next = l2
//	if tmp1 == nil {
//		l1.Next.Next = tmp2
//	} else {
//		l1.Next.Next = tmp1
//	}
//	l1, l2 = tmp1, tmp2
//}
//}

func mergeListFromLeft(l1, l2 *s.ListNode) *s.ListNode {
	dummy := new(s.ListNode)
	head, toggle := dummy, true
	for l1 != nil && l2 != nil {
		// 节点切换
		if toggle {
			head.Next, l1 = l1, l1.Next
			//head.Next = l1
			//head, l1 = head.Next, l1.Next
		} else {
			head.Next, l2 = l2, l2.Next
		}
		head, toggle = head.Next, !toggle
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return dummy.Next
}
