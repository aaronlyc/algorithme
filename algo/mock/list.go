package mock

import s "example.com/algorithms/algo/struct"

func MackIntSlice2List(intSlice []int) *s.ListNode {
	dummy := new(s.ListNode)
	head := dummy
	for i, val := range intSlice {
		head.Val = val
		if i < len(intSlice)-1 {
			head.Next = new(s.ListNode)
			head = head.Next
		}
	}
	return dummy
}
