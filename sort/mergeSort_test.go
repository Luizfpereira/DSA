package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	A := []int{2, 8, 7, 5, 10, 8, 1}
	want := []int{1, 2, 5, 7, 8, 8, 10}

	got := mergeSort(A)

	assert.Equal(t, want, got)
}

func TestMergeSortIdx(t *testing.T) {
	t.Run("test merge in entire slice", func(t *testing.T) {
		A := []int{2, 8, 7, 5, 10, 8, 1}
		want := []int{1, 2, 5, 7, 8, 8, 10}

		mergeSortIdx(A, 0, len(A)-1)

		assert.Equal(t, want, A)
	})
}
