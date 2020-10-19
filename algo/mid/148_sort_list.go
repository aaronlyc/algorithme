package mid

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import (
	s "example.com/algorithms/algo/struct"
)

// 方法一: 递归
// 时间复杂度：O(nlogn)
// 空间复杂度: O(logn)
func sortList(head *s.ListNode) *s.ListNode {
	// 递归结束条件是: 链表为空或者链表只含一个节点
	if head == nil || head.Next == nil {
		return head
	}
	// 1.先分传入数据
	// 需要一个可以找到链表中间节点的函数, 用来每次对半拆分链表
	midNode := findMidNode(head)
	rightList := midNode.Next
	midNode.Next = nil
	leftSortList, rigthSortList := sortList(head), sortList(rightList)

	// 2. 后和两个排序的链表
	return mergeTwoSortList(leftSortList, rigthSortList)
}

// 快慢指针的方式
func findMidNode(head *s.ListNode) *s.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	// 结束条件: fast == nil || fast.Next == nil
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func mergeTwoSortList(l1, l2 *s.ListNode) *s.ListNode {
	dummy := new(s.ListNode)
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next, l1 = l1, l1.Next
		} else {
			head.Next, l2 = l2, l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return dummy.Next
}

// 方法二: 迭代
// 时间复杂度：O(nlogn)
// 空间复杂度: O(1)
