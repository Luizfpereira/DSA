package linkedlist

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

type LinkedList[T any] struct {
	Size int
	Head *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{Val: val, Next: nil}
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) addFirst(val T) {
	l.Size++
	newNode := NewNode(val)
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *LinkedList[T]) addLast(val T) {
	l.Size++
	newNode := NewNode(val)
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *LinkedList[T]) deleteFirst() (T, bool) {
	if l.Head != nil {
		curr := l.Head
		l.Head = curr.Next
		l.Size--
		return curr.Val, true
	}

	var r T
	return r, false
}

func (l *LinkedList[T]) deleteLast() (T, bool) {

	if l.Head == nil {
		var r T
		return r, false
	}

	if l.Head.Next == nil {
		return l.deleteFirst()
	}

	curr := l.Head

	for ; curr.Next.Next != nil; curr = curr.Next {
	}

	retval := curr.Next.Val
	curr.Next = nil
	l.Size--

	return retval, true
}

// func (l *LinkedList[T]) findNodeAt()
