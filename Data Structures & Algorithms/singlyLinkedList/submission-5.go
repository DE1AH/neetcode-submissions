type Node struct{
	Next *Node
	Val int
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	curr := ll.head
	for i := 0; i < index && curr != nil; i++{
		curr = curr.Next
	}
	if curr == nil{
		return -1
	}
	return curr.Val
}

func (ll *LinkedList) InsertHead(val int) {
	nn := &Node{Val : val}
	if ll.head == nil{
		ll.head = nn
		return
	}
	nn = &Node{Next : ll.head, Val : val}
	ll.head = nn
}

func (ll *LinkedList) InsertTail(val int) {
	nn := &Node{Next:nil,Val : val}
	if ll.head == nil{
		ll.head = nn
		return
	}
	curr := ll.head
	for curr.Next != nil{
		curr = curr.Next
	}
	curr.Next = nn
}

func (ll *LinkedList) Remove(index int) bool {
	if ll.head == nil {
		return false
	}
	if index == 0{
		ll.head = ll.head.Next
		return true
	}
	curr := ll.head
	for i := 0; i < index - 1 && curr != nil; i++{
		curr = curr.Next
	}
	if curr == nil || curr.Next == nil{
		return false
	}
	curr.Next = curr.Next.Next
	return true
}

func (ll *LinkedList) GetValues() []int {
	arr := make([]int,0)
	curr := ll.head
	for curr != nil{
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	return arr
}
