/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l *ListNode, r *ListNode) *ListNode{
	if l == nil{
		return r
	}
	if r == nil{
		return l
	}
	dummy := &ListNode{}
	tail := dummy
	for l != nil && r != nil{
		if l.Val <= r.Val{
			tail.Next = l
			l = l.Next
		} else{
			tail.Next = r
			r = r.Next
		}
		tail = tail.Next
	}

	if l != nil{
		tail.Next = l
	}
	if r != nil{
		tail.Next = r
	}

	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0{
		return nil
	}
	interval := 1
	for interval < len(lists){
		for i := 0; i < len(lists) - interval; i += 2*interval{
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
			
		}
		interval *= 2
	}
	return lists[0]
}
