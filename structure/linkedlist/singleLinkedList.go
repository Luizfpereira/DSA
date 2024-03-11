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
	newNode.Next = l.Head
	l.Head = newNode
}
