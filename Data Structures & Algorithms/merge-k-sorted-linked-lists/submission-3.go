/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode{
	dummy := &ListNode{}
	tail := dummy
	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil{
		tail.Next = l1
	}
	if l2 != nil{
		tail.Next = l2
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	interval := 1
	for interval < len(lists){
		for i := 0; i < len(lists) - interval; i += 2 * interval{
			lists[i] = mergeTwoLists(lists[i], lists[i + interval])
		}
		interval *= 2
	}
	return lists[0]
}
