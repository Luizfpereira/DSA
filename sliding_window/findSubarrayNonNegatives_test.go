package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubArraySumNestedLoop(t *testing.T) {
	t.Run("first case", func(t *testing.T) {
		arr := []int{1, 4, 20, 3, 10, 5}
		sum := 33

		exists := subArraySumNestedLoop(arr, sum)

		assert.True(t, exists)
	})

	t.Run("second case", func(t *testing.T) {
		arr := []int{1, 4, 0, 0, 3, 10, 5}
		sum := 7

		exists := subArraySumNestedLoop(arr, sum)

		assert.True(t, exists)
	})

	t.Run("third case", func(t *testing.T) {
		arr := []int{1, 4}
		sum := 0

		exists := subArraySumNestedLoop(arr, sum)

		assert.False(t, exists)
	})
}

func TestSubArraySumSlidingWindow(t *testing.T) {
	t.Run("first case", func(t *testing.T) {
		arr := []int{1, 4, 20, 3, 10, 5}
		sum := 33

		exists := subArraySumSlidingWindow(arr, sum)

		assert.True(t, exists)
	})

	t.Run("second case", func(t *testing.T) {
		arr := []int{1, 4, 0, 0, 3, 10, 5}
		sum := 7

		exists := subArraySumSlidingWindow(arr, sum)

		assert.True(t, exists)
	})

	t.Run("third case", func(t *testing.T) {
		arr := []int{1, 4}
		sum := 0

		exists := subArraySumSlidingWindow(arr, sum)

		assert.False(t, exists)
	})
}
