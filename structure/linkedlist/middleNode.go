package linkedlist

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func (ll *LinkedList) GetMiddleNode() *LinkedList {
	depth := ll.countDepth()

	i := 0

	for i < depth/2 { 
		if ll.Next == nil {
			return ll
		}
		ll = ll.Next
		i++
	}
	return ll
}

func (ll *LinkedList) countDepth() int {
	depth := 0

	for ll != nil {
		depth++
		if ll.Next == nil {
			break
		}
		ll = ll.Next
	}
	return depth
}
