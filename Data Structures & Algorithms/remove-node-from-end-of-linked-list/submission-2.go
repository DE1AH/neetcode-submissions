/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next : head}
	s,f := dummy, head
	for n > 0{
		f = f.Next
		n --
	}
	for f != nil {
		s = s.Next
		f = f.Next
	}
	s.Next = s.Next.Next
	return dummy.Next
}
