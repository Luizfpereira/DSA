package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMiddleNode(t *testing.T) {
	t.Run("even case", func(t *testing.T) {
		arr := []int{2, 7, 3, 5}

		aux := &LinkedList{Val: arr[0]}
		ll := aux
		for _, a := range arr[1:] {
			aux.Next = &LinkedList{Val: a}
			aux = aux.Next
		}

		want := &LinkedList{Val: 3}
		want.Next = &LinkedList{Val: 5}

		got := ll.GetMiddleNode()

		assert.Equal(t, want, got)
	})

	t.Run("odd case", func(t *testing.T) {

		arr := []int{2, 7, 3, 5, 9}

		aux := &LinkedList{Val: arr[0]}
		ll := aux
		for _, a := range arr[1:] {
			aux.Next = &LinkedList{Val: a}
			aux = aux.Next
		}

		aux = &LinkedList{Val: arr[2]}
		want := aux
		for _, a := range arr[3:] {
			aux.Next = &LinkedList{Val: a}
			aux = aux.Next
		}

		got := ll.GetMiddleNode()

		assert.Equal(t, want, got)
	})
}
