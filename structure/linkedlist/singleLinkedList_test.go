package linkedlist

import (
	"fmt"
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

func TestDeleteFirst(t *testing.T) {
	t.Run("List with 2 items", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.addFirst(1)
		l.addFirst(2)

		val, deleted := l.deleteFirst()
		current := l.Head

		want := []int{1}
		var got []int
		got = append(got, current.Val)

		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}

		assert.Equal(t, want, got)
		assert.Equal(t, 1, l.Size)
		assert.Equal(t, 2, val)
		assert.Equal(t, true, deleted)
	})

	t.Run("List with 1 item1", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.addFirst(1)

		val, deleted := l.deleteFirst()

		assert.Nil(t, l.Head)
		assert.Equal(t, 0, l.Size)
		assert.Equal(t, 1, val)
		assert.Equal(t, true, deleted)
	})

	t.Run("Empty list", func(t *testing.T) {
		l := NewLinkedList[int]()
		val, deleted := l.deleteFirst()
		assert.Equal(t, 0, val)
		assert.Equal(t, false, deleted)
	})
}

func TestDeleteLast(t *testing.T) {
	t.Run("List with 2 items", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.addFirst(1)
		l.addFirst(2)

		val, deleted := l.deleteLast()
		current := l.Head

		want := []int{2}
		var got []int
		got = append(got, current.Val)

		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}

		assert.Equal(t, want, got)
		assert.Equal(t, 1, l.Size)
		assert.Equal(t, 1, val)
		assert.Equal(t, true, deleted)
	})

	t.Run("List with 1 item1", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.addFirst(2)

		val, deleted := l.deleteLast()

		assert.Nil(t, l.Head)
		assert.Equal(t, 0, l.Size)
		assert.Equal(t, 2, val)
		assert.Equal(t, true, deleted)
	})

	t.Run("Empty list", func(t *testing.T) {
		l := NewLinkedList[int]()
		val, deleted := l.deleteLast()
		assert.Equal(t, 0, val)
		assert.Equal(t, false, deleted)
	})
}

func TestFindNodeAt(t *testing.T) {
	l := NewLinkedList[int]()
	l.addFirst(1)
	l.addFirst(7)
	l.addFirst(11)

	// 11, 7, 1

	fmt.Println(l.Head, l.Head.Val)
	fmt.Println(l.Head.Next, l.Head.Next.Val)
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
