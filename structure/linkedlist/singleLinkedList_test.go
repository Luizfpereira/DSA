package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFirst(t *testing.T) {
	l := NewLinkedList[int]()
	l.addFirst(1)
	l.addFirst(5)
	l.addFirst(7)

	want := []int{7, 5, 1}
	var got []int

	current := l.Head

	got = append(got, current.Val)

	for current.Next != nil {
		current = current.Next
		got = append(got, current.Val)
	}

	assert.EqualValues(t, want, got)
}

func TestAddLast(t *testing.T) {
	t.Run("test empty linkedList", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.addLast(1)

		want := []int{1}
		var got []int

		current := l.Head

		got = append(got, current.Val)
		assert.EqualValues(t, want, got)
	})

	t.Run("test not empty linkedList", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.addFirst(1)
		l.addFirst(5)
		l.addFirst(7)
		l.addLast(9)
		l.addLast(3)

		//7, 5, 1, 9, 3

		want := []int{7, 5, 1, 9, 3}
		var got []int

		current := l.Head

		got = append(got, current.Val)

		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}

		assert.EqualValues(t, want, got)
	})
}

//addFirst
//addLast
//addAfterValue
//addBeforeValue
//DeleteFirst
//DeleteEnd
//DeleteAt
//findNodeAt
//getFirst
//getLast
//toArray
//size
//reverse
//reversePartition
//print
